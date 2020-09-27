package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/ichi-pg/golang-server/internal/domain"
)

const paymentLogKind = "PaymentLog"

type paymentRepository struct {
	masterRepo domain.MasterRepository
}

// PaymentRepository はPaymentRepositoryのDatastore実装を返します。
func PaymentRepository(masterRepo domain.MasterRepository) domain.PaymentRepository {
	return paymentRepository{
		masterRepo: masterRepo,
	}
}

func (r paymentRepository) Pay(c context.Context, user *domain.User, paymentItem *domain.PaymentItem) (*domain.PaymentLog, error) {
	log := domain.PaymentLog{
		Item:      *paymentItem,
		CreatedAt: time.Now(),
	}
	err := runWithClient(c, func(cli *datastore.Client) error {
		key := itemInventoryKey(user.ID, paymentItem.Item.ID)
		inventory := domain.ItemInventory{
			UserID:   user.ID,
			ItemID:   paymentItem.Item.ID,
			Quantity: 0,
		}
		err := cli.Get(c, key, &inventory)
		if err != datastore.ErrNoSuchEntity {
			return err
		}
		inventory.Quantity += paymentItem.Quantity
		_, err = cli.Put(c, key, &inventory)
		if err != nil {
			return err
		}
		_, err = cli.Put(c, newKey(paymentLogKind, log.CreatedAt.String(), userKey(user.ID)), &log)
		return err
	})
	return &log, err
}

func (r paymentRepository) Logs(c context.Context, user *domain.User, cursor domain.Cursor) ([]domain.PaymentLog, domain.Cursor, error) {
	logs := []domain.PaymentLog{}
	nextCursor, err := runPagination(c, newQuery(paymentLogKind).Ancestor(userKey(user.ID)), 20, cursor, func(it *datastore.Iterator) error {
		log := domain.PaymentLog{}
		_, err := it.Next(&log)
		if err != nil {
			return err
		}
		item, err := r.masterRepo.PaymentItem(c, log.Item.ID)
		if err != nil {
			return err
		}
		log.Item = *item
		logs = append(logs, log)
		return nil
	})
	return logs, nextCursor, err
}
