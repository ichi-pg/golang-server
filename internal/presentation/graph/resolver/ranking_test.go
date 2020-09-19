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

func TestQuery_Rankers(t *testing.T) {
	t.Parallel()

	c := context.Background()
	r := NewResolver(logrus.NewEntry(logrus.New()), nil, injection.MockInjector())

	res, err := r.Query().Rankers(c, 0, 10)
	if assert.NoError(t, err) {
		assert.Equal(t, res, []*generated.Ranker{
			{
				User:  newUser(&mock.User),
				Rank:  1,
				Score: 120,
			},
			{
				User:  newUser(&mock.RankingUserA),
				Rank:  2,
				Score: 100,
			},
			{
				User:  newUser(&mock.RankingUserB),
				Rank:  3,
				Score: 80,
			},
		})
	}
}
