package presentation

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/vectorman1/product-search-proto/product_service"

	"github.com/elastic/go-elasticsearch/v8"
)

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
		doc := hit["_source"]
		documentBytes, err := json.Marshal(doc)
		if err != nil {
			return nil, err
		}

		product := new(product_service.Product)
		err = json.Unmarshal(documentBytes, product)
		if err != nil {
			return nil, err
		}
		result = append(result, product)
	}

	return &product_service.SearchResponse{
		Items: result,
	}, nil
}
