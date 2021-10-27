package presentation

import (
	"context"
	"product_service/generated/product_service"

	"github.com/elastic/go-elasticsearch/v8"
)

type productServiceServer struct {
	es *elasticsearch.Client
	product_service.UnimplementedProductServiceServer
}

func NewProductServiceServer(elasticClient *elasticsearch.Client) *productServiceServer {
	return &productServiceServer{es: elasticClient}
}

func (s *productServiceServer) Search(ctx context.Context,req *product_service.ProductSearchRequest) (*product_service.ProductSearchResponse, error) {
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query":
			},
		},
	}
	panic("unimplemented")
}
