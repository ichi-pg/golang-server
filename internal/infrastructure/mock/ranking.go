package mock

import (
	"context"

	"github.com/ichi-pg/golang-server/internal/domain"
)

type rankingRepository struct {
}

// RankingRepository はRankingRepositoryのモック実装を返します。
func RankingRepository() domain.RankingRepository {
	return rankingRepository{}
}

func (r rankingRepository) Rankers(c context.Context, offset, limit int64) ([]domain.Ranker, error) {
	return []domain.Ranker{
		{
			UserID: UserID,
			Rank:   1,
			Score:  120,
		},
		{
			UserID: domain.UserID("aaaa"),
			Rank:   2,
			Score:  100,
		},
		{
			UserID: domain.UserID("bbbb"),
			Rank:   3,
			Score:  80,
		},
	}, nil
}
