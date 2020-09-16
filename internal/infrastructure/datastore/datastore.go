package datastore

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
	submoduleEnv "github.com/ichi-pg/golang-middleware/env"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/pkg/env"
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

func newQuery(kind string) *datastore.Query {
	return datastore.NewQuery(kind).Namespace(os.Getenv(env.Namespace))
}

func newKey(kind, name string, parent *datastore.Key) *datastore.Key {
	key := datastore.NameKey(userKind, name, parent)
	key.Namespace = os.Getenv(env.Namespace)
	return key
}
