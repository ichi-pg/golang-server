package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ichi-pg/golang-server/internal/presentation/graph/generated"
)

func (r *queryResolver) User(ctx context.Context) (*generated.User, error) {
	return newUser(r.Resolver.User), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
