package presentation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/vectorman1/product-search/grpc/generated/product_service"

	"github.com/elastic/go-elasticsearch/v8"
)

type searchResponse struct {
	Hits struct {
		Hits []struct {
			product_service.Product `json:"_source"`
		}
	}
}

type productServiceServer struct {
	es    *elasticsearch.Client
	index string
	product_service.UnimplementedProductServiceServer
}

func NewProductServiceServer(es *elasticsearch.Client) *productServiceServer {
	return &productServiceServer{
		es:    es,
		index: "index_products",
	}
}

func (s *productServiceServer) Search(ctx context.Context, req *product_service.SearchRequest) (*product_service.SearchResponse, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"size": req.GetPageSize(),
		"_source": map[string]interface{}{
			"includes": req.GetFields(),
		},
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": req.GetQuery(),
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := s.es.Search(
		s.es.Search.WithContext(ctx),
		s.es.Search.WithIndex(s.index),
		s.es.Search.WithBody(&buf),
		s.es.Search.WithFilterPath("hits.hits._source"))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	result := make([]*product_service.Product, 0, req.GetPageSize())
	hits := response["hits"].(map[string]interface{})["hits"].([]map[string]interface{})
	for _, hit := range hits {
		product, ok := hit["_source"].(*product_service.Product)
		if !ok {
			return nil, fmt.Errorf("data is not a product")
		}
		result = append(result, product)
	}

	return &product_service.SearchResponse{
		Items: result,
	}, nil
}
