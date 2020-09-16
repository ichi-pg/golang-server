package resolver

import (
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/presentation/graph/generated"
)

func newRanker(ranker *domain.Ranker) *generated.Ranker {
	if ranker == nil {
		return nil
	}
	return &generated.Ranker{
		UserID: string(ranker.UserID),
		Rank:   ranker.Rank,
		Score:  ranker.Score,
	}
}

func newRankers(rankers []domain.Ranker) []*generated.Ranker {
	res := make([]*generated.Ranker, len(rankers))
	for i, ranker := range rankers {
		res[i] = newRanker(&ranker)
	}
	return res
}
