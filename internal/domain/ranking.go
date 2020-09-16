package domain

import (
	"context"
)

// Ranker はランキング内のユーザーです。
type Ranker struct {
	UserID
	Rank  int64
	Score int64
}

// RankingRepository はランキングのCRUDを抽象化します。
type RankingRepository interface {
	Rankers(c context.Context, offset, limit int64) ([]*Ranker, error)
}
