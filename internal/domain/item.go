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
