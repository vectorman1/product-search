package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/vectorman1/product-search-proto/product_service"
	"github.com/vectorman1/product-search/graphql-api/graph/generated"
	"github.com/vectorman1/product-search/graphql-api/graph/model"
)

func (r *queryResolver) Search(ctx context.Context, query string, pageSize int) ([]*model.Product, error) {
	req := &product_service.SearchRequest{
		Query:    query,
		PageSize: int32(pageSize),
	}

	// get context of resolver
	resolverContext := ctx.Value("resolver_context").(graphql.FieldContext)

	// get fields selected from query to pass to elastic
	// more complex logic would be necessary for nested objects
	selectedFields := resolverContext.Field.Selections
	selectedFieldsLen := len(selectedFields)
	fields := make([]string, selectedFieldsLen, selectedFieldsLen)
	for _, field := range selectedFields {
		fields = append(fields, field.(*ast.Field).Name)
	}
	req.Fields = fields

	// perform the search request to grpc service
	res, err := r.productServiceClient.Search(ctx, req)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Product, 0, len(res.GetItems()))
	for _, item := range res.GetItems() {
		result = append(result, &model.Product{
			Category:            item.Category,
			Subcategory:         item.Subcategory,
			Name:                item.Name,
			CurrentPrice:        float64(item.CurrentPrice),
			RawPrice:            float64(item.RawPrice),
			Currency:            item.Currency,
			Discount:            int(item.Discount),
			LikesCount:          int(item.LikesCount),
			IsNew:               item.IsNew,
			Brand:               item.Brand,
			BrandURL:            item.BrandUrl,
			CodCountry:          item.CodCountry,
			Variation0Color:     item.Variation0Color,
			Variation1Color:     item.Variation1Color,
			Variation0Thumbnail: item.Variation0Thumbnail,
			Variation0Image:     item.Variation0Image,
			Variation1Thumbnail: item.Variation1Thumbnail,
			Variation1Image:     item.Variation1Image,
			ImageURL:            item.ImageUrl,
			URL:                 item.Url,
			ID:                  int(item.Id),
			Model:               item.Model,
		})
	}

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
