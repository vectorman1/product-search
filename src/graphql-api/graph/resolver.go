package graph

import "github.com/vectorman1/product-search-proto/product_service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productServiceClient product_service.ProductServiceClient
}
