package resolver

import (
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/presentation/graph/generated"
)

func newUser(user *domain.User) *generated.User {
	if user == nil {
		return nil
	}
	return &generated.User{
		UserID:    string(user.ID),
		Name:      string(user.Name),
		CreatedAt: user.CreatedAt,
	}
}
