package domain

import (
	"context"
	"time"
)

type (
	// PaymentItemID は課金アイテムのIDです。
	PaymentItemID string
)

// PaymentItem は課金アイテムの値段と個数です。
type PaymentItem struct {
	Item     Item
	Price    int64
	Quantity int64
}

// PaymentLog は課金履歴です。
type PaymentLog struct {
	Item      PaymentItem
	CreatedAt time.Time
}

// PaymentRepository は課金のCRUDを抽象化します。
type PaymentRepository interface {
	Pay(c context.Context, user *User, paymentItemID PaymentItemID) (*PaymentLog, error)
	Logs(c context.Context, user *User, cursor Cursor) ([]PaymentLog, Cursor, error)
}
