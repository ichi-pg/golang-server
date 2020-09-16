package redis

import (
	"context"
	"math"

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

func (r rankingRepository) add(c context.Context, userID domain.UserID, score int64) error {
	return runWithClient(func(cli *redis.Client) error {
		return cli.ZAdd(ranking, redis.Z{
			Member: string(userID),
			Score:  float64(score),
		}).Err()
	})
}

func (r rankingRepository) Rankers(c context.Context, offset, limit int64) ([]domain.Ranker, error) {
	rankers := []domain.Ranker{}
	err := runWithClient(func(cli *redis.Client) error {
		res, err := cli.ZRevRangeWithScores(ranking, offset, offset+limit-1).Result()
		if err != nil {
			return err
		}
		prevRank := offset + 1
		prevScore := int64(math.MinInt64)
		for i, v := range res {
			var rank int64
			score := int64(v.Score)
			if prevScore == score {
				rank = prevRank
			} else {
				rank = offset + int64(i) + 1
				prevRank = rank
				prevScore = score
			}
			rankers = append(rankers, domain.Ranker{
				UserID: domain.UserID(v.Member.(string)),
				Rank:   rank,
				Score:  score,
			})
		}
		return nil
	})
	return rankers, err
}
