package redis

import (
	"context"
	"testing"

	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/infrastructure/mock"
	"github.com/stretchr/testify/assert"
)

func TestRanking_Rankers(t *testing.T) {
	t.Parallel()

	c := context.Background()
	r := rankingRepository{
		mock.UserRepository(),
	}

	assert.NoError(t, r.add(c, mock.User.ID, 100))
	assert.NoError(t, r.add(c, mock.User.ID, 110))
	assert.NoError(t, r.add(c, mock.RankingUserA.ID, 80))
	assert.NoError(t, r.add(c, mock.RankingUserB.ID, 120))
	assert.NoError(t, r.add(c, mock.RankingUserC.ID, 120))

	rankers, err := r.Rankers(c, 0, 10)
	if assert.NoError(t, err) {
		assert.Equal(t, rankers, []domain.Ranker{
			{
				User:  mock.RankingUserB,
				Rank:  1,
				Score: 120,
			},
			{
				User:  mock.RankingUserC,
				Rank:  1,
				Score: 120,
			},
			{
				User:  mock.User,
				Rank:  3,
				Score: 110,
			},
			{
				User:  mock.RankingUserA,
				Rank:  4,
				Score: 80,
			},
		})
	}
}
