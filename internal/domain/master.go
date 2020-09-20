package domain

import (
	"context"
)

// MasterRepository はマスターデータの読み取りを抽象化します。
type MasterRepository interface {
	PaymentItems(c context.Context) ([]PaymentItem, error)
	PaymentItem(c context.Context, id PaymentItemID) (*PaymentItem, error)
}
