package redis

import (
	"context"
	"math"

	"github.com/go-redis/redis"
	"github.com/ichi-pg/golang-server/internal/domain"
)

const rankingKey = "ranking"

type rankingRepository struct {
	userRepo domain.UserRepository
}

// RankingRepository はRankingRepositoryのRedis実装を返します。
func RankingRepository(userRepo domain.UserRepository) domain.RankingRepository {
	return rankingRepository{
		userRepo: userRepo,
	}
}

func (r rankingRepository) add(c context.Context, userID domain.UserID, score int64) error {
	return runWithClient(func(cli *redis.Client) error {
		return cli.ZAdd(rankingKey, redis.Z{
			Member: string(userID),
			Score:  float64(score),
		}).Err()
	})
}

func (r rankingRepository) Rankers(c context.Context, offset, limit int64) ([]domain.Ranker, error) {
	rankers := []domain.Ranker{}
	err := runWithClient(func(cli *redis.Client) error {
		res, err := cli.ZRevRangeWithScores(rankingKey, offset, offset+limit-1).Result()
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
			user, err := r.userRepo.ByUserID(c, domain.UserID(v.Member.(string)))
			if err != nil {
				return err
			}
			rankers = append(rankers, domain.Ranker{
				User:  *user,
				Rank:  rank,
				Score: score,
			})
		}
		return nil
	})
	return rankers, err
}
