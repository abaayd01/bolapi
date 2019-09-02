package resolvers

import (
	"bolapi/internal/pkg/gql"
	"context"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (q queryResolver) Hello(ctx context.Context) (*string, error) {
	str := "world!"
	return &str, nil
}
