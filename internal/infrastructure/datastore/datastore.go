package datastore

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
	submoduleEnv "github.com/ichi-pg/golang-middleware/env"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/pkg/env"
	"google.golang.org/api/iterator"
)

func runWithClient(c context.Context, f func(*datastore.Client) error) error {
	cli, err := datastore.NewClient(c, os.Getenv(submoduleEnv.ProjectID))
	if err != nil {
		return err
	}
	defer cli.Close()
	err = f(cli)
	if err == datastore.ErrNoSuchEntity {
		return domain.ErrNoSuchEntity
	}
	return err
}

func runWithTransaction(c context.Context, f func(*datastore.Transaction) error) error {
	return runWithClient(c, func(cli *datastore.Client) error {
		tx, err := cli.NewTransaction(c)
		if err != nil {
			return err
		}
		err = f(tx)
		if err != nil {
			e := tx.Rollback()
			if e != nil {
				return e
			}
			return err
		}
		_, err = tx.Commit()
		return err
	})
}

func newQuery(kind string) *datastore.Query {
	return datastore.NewQuery(kind).Namespace(os.Getenv(env.Namespace))
}

func newKey(kind, name string, parent *datastore.Key) *datastore.Key {
	key := datastore.NameKey(userKind, name, parent)
	key.Namespace = os.Getenv(env.Namespace)
	return key
}

func runPagination(c context.Context, q1 *datastore.Query, limit int, cursor domain.Cursor, next func(*datastore.Iterator) error) (domain.Cursor, error) {
	q2 := q1.Limit(limit)
	if !cursor.Empty() {
		cursor, err := datastore.DecodeCursor(string(cursor))
		if err != nil {
			return "", err
		}
		q2 = q2.Start(cursor)
	}
	var nextCursor domain.Cursor
	err := runWithClient(c, func(cli *datastore.Client) error {
		it := cli.Run(c, q2)
		var err error
		for err == nil {
			err = next(it)
		}
		if err != iterator.Done {
			return err
		}
		cur, err := it.Cursor()
		if err != nil {
			return err
		}
		nextCursor = domain.Cursor(cur.String())
		return nil
	})
	return nextCursor, err
}
