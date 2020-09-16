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
				UserID: string(mock.UserID),
				Rank:   1,
				Score:  120,
			},
			{
				UserID: "aaaa",
				Rank:   2,
				Score:  100,
			},
			{
				UserID: "bbbb",
				Rank:   3,
				Score:  80,
			},
		})
	}
}
