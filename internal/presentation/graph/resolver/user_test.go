package resolver

import (
	"context"
	"testing"

	"github.com/ichi-pg/golang-server/internal/application/injection"
	"github.com/ichi-pg/golang-server/internal/infrastructure/mock"
	"github.com/ichi-pg/golang-server/internal/presentation/graph/generated"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestQuery_User(t *testing.T) {
	t.Parallel()

	mockUser := mock.NewUser()
	c := context.Background()
	r := NewResolver(logrus.NewEntry(logrus.New()), mockUser, injection.MockInjector())

	res, err := r.Query().User(c)
	if assert.NoError(t, err) {
		assert.Equal(t, &generated.User{
			UserID:    string(mock.UserID),
			Name:      string(mock.UserName),
			CreatedAt: mock.UserCreateAt,
		}, res)
	}
}
