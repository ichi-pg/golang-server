package application

import (
	"github.com/ichi-pg/golang-server/internal/domain"
)

// RankingUsecase はランキングのユースケースを実装します。
type RankingUsecase struct {
	repo domain.RankingRepository
}

// NewRankingUsecase はランキングのユースケースを生成します。
func NewRankingUsecase(repo domain.RankingRepository) RankingUsecase {
	return RankingUsecase{
		repo: repo,
	}
}

// Rankers はランキングのユーザーリストを返します。
func (u RankingUsecase) Rankers(c AuthContext, offset, limit int64) ([]domain.Ranker, error) {
	return u.repo.Rankers(c.Context, offset, limit)
}
