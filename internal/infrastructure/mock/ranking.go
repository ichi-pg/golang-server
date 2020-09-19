package mock

import (
	"context"

	"github.com/ichi-pg/golang-server/internal/domain"
)

// RankingUserA はモックにおける正常系のユーザーです。
var RankingUserA = domain.User{
	ID:   domain.UserID("e71220ec-2183-4130-a1a5-ce5cd9bc4d01"),
	Name: domain.UserName("aaaa"),
}

// RankingUserB はモックにおける正常系のユーザーです。
var RankingUserB = domain.User{
	ID:   domain.UserID("6d239877-1031-4fc0-bb9d-b5acc3779f61"),
	Name: domain.UserName("bbbb"),
}

// RankingUserC はモックにおける正常系のユーザーです。
var RankingUserC = domain.User{
	ID:   domain.UserID("3c7d7f9d-de22-40c7-9eeb-97e6cd0b2378"),
	Name: domain.UserName("cccc"),
}

type rankingRepository struct {
}

// RankingRepository はRankingRepositoryのモック実装を返します。
func RankingRepository() domain.RankingRepository {
	return rankingRepository{}
}

func (r rankingRepository) Rankers(c context.Context, offset, limit int64) ([]domain.Ranker, error) {
	return []domain.Ranker{
		{
			User:  User,
			Rank:  1,
			Score: 120,
		},
		{
			User:  RankingUserA,
			Rank:  2,
			Score: 100,
		},
		{
			User:  RankingUserB,
			Rank:  3,
			Score: 80,
		},
	}, nil
}
