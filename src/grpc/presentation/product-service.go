package presentation

import (
	"context"
	"product_service/generated/product_service"
)

type productServiceServer struct {
	product_service.UnimplementedProductServiceServer
}

func NewProductServiceServer() *productServiceServer {
	return &productServiceServer{}
}

func (s *productServiceServer) Search(context.Context, *product_service.ProductSearchRequest) (*product_service.ProductSearchResponse, error) {
	panic("unimplemented")
}
