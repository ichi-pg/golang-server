package datastore

import (
	"os"

	"cloud.google.com/go/datastore"
	"github.com/ichi-pg/golang-server/internal/pkg/env"
)

func newQuery(kind string) *datastore.Query {
	return datastore.NewQuery(kind).Namespace(os.Getenv(env.Namespace))
}

func newKey(kind, name string, parent *datastore.Key) *datastore.Key {
	key := datastore.NameKey(userKind, name, parent)
	key.Namespace = os.Getenv(env.Namespace)
	return key
}
