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
	r := rankingRepository{}

	assert.NoError(t, r.add(c, mock.UserID, 100))
	assert.NoError(t, r.add(c, mock.UserID, 110))
	assert.NoError(t, r.add(c, domain.UserID("aaaa"), 80))
	assert.NoError(t, r.add(c, domain.UserID("bbbb"), 120))
	assert.NoError(t, r.add(c, domain.UserID("cccc"), 120))

	rankers, err := r.Rankers(c, 0, 10)
	if assert.NoError(t, err) {
		assert.Equal(t, rankers, []domain.Ranker{
			{
				UserID: domain.UserID("cccc"),
				Rank:   1,
				Score:  120,
			},
			{
				UserID: domain.UserID("bbbb"),
				Rank:   1,
				Score:  120,
			},
			{
				UserID: mock.UserID,
				Rank:   3,
				Score:  110,
			},
			{
				UserID: domain.UserID("aaaa"),
				Rank:   4,
				Score:  80,
			},
		})
	}
}
