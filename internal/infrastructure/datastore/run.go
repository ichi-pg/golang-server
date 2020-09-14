package datastore

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/ichi-pg/golang-middleware/env"
	"github.com/ichi-pg/golang-server/internal/domain"
)

func runWithClient(c context.Context, f func(*datastore.Client) error) error {
	cli, err := datastore.NewClient(c, os.Getenv(env.ProjectID))
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
