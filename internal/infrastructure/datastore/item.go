package datastore

import (
	"cloud.google.com/go/datastore"
	"github.com/ichi-pg/golang-server/internal/domain"
)

const itemInventoryKind = "ItemInventory"

func itemInventoryKey(userID domain.UserID, itemID domain.ItemID) *datastore.Key {
	return newKey(itemInventoryKind, string(itemID), userKey(userID))
}
