package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ichi-pg/golang-server/internal/application"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/presentation/graph/generated"
)

func (r *mutationResolver) UpdateUserName(ctx context.Context, name string) (*generated.User, error) {
	user, err := r.Injector.UserUsecase().UpdateName(application.NewUserContext(ctx, r.Logger, r.User), domain.UserName(name))
	return newUser(user), err
}

func (r *mutationResolver) Pay(ctx context.Context, itemID string) (*generated.PaymentLog, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
