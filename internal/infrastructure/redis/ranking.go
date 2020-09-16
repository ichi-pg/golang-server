package redis

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/ichi-pg/golang-server/internal/domain"
)

const ranking = "ranking"

type rankingRepository struct {
}

// RankingRepository はRankingRepositoryのRedis実装を返します。
func RankingRepository() domain.RankingRepository {
	return rankingRepository{}
}

func (r rankingRepository) Rankers(c context.Context, offset, limit int64) ([]*domain.Ranker, error) {
	rankers := []*domain.Ranker{}
	err := runWithClient(func(cli *redis.Client) error {
		res, err := cli.ZRangeWithScores(ranking, offset, offset+limit).Result()
		if err != nil {
			return err
		}
		for i, v := range res {
			rankers = append(rankers, &domain.Ranker{
				UserID: domain.UserID(v.Member.(string)),
				Rank:   offset + int64(i),
				Score:  int64(v.Score),
			})
		}
		return nil
	})
	return rankers, err
}

//TODO test code
