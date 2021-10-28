package graph

import "github.com/vectorman1/product-search/grpc/generated/product_service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	product_service.ProductServiceClient
}
