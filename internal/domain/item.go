package domain

type (
	// ItemID はアイテムのIDです。
	ItemID string
)

// Item はアイテムのマスターデータです。
type Item struct {
	ID   ItemID
	Name string
}

// ItemInventory はアイテムの所持数です。
type ItemInventory struct {
	UserID   UserID
	ItemID   ItemID
	Quantity int64
}
