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

	c := context.Background()
	r := NewResolver(logrus.NewEntry(logrus.New()), &mock.User, injection.MockInjector())

	res, err := r.Query().User(c)
	if assert.NoError(t, err) {
		assert.Equal(t, res, &generated.User{
			ID:   string(mock.User.ID),
			Name: string(mock.User.Name),
		})
	}
}
