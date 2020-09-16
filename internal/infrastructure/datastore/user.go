package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/ichi-pg/golang-server/internal/domain"
)

const userKind = "User"

type userRepository struct {
}

// UserRepository はUserRepositoryのDatastore実装を返します。
func UserRepository() domain.UserRepository {
	return userRepository{}
}

func (r userRepository) ByUserID(c context.Context, userID domain.UserID) (*domain.User, error) {
	key := newKey(userKind, string(userID), nil)
	user := &domain.User{}
	err := runWithClient(c, func(cli *datastore.Client) error {
		return cli.Get(c, key, user)
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r userRepository) ByDummyID(c context.Context, dummyID domain.DummyID) (*domain.User, error) {
	q := newQuery(userKind).Filter("DummyID =", string(dummyID))
	return r.byQuery(c, q)
}

func (r userRepository) ByFirebaseID(c context.Context, firebaseID domain.FirebaseID) (*domain.User, error) {
	q := newQuery(userKind).Filter("FirebaseID =", string(firebaseID))
	return r.byQuery(c, q)
}

func (r userRepository) byQuery(c context.Context, q *datastore.Query) (*domain.User, error) {
	users := []*domain.User{}
	err := runWithClient(c, func(cli *datastore.Client) error {
		_, err := cli.GetAll(c, q, &users)
		return err
	})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, domain.ErrNoSuchEntity
	}
	return users[0], nil
}

func (r userRepository) CreateDummyUser(c context.Context, dummyID domain.DummyID) (*domain.User, error) {
	user := domain.NewDummyUser(dummyID)
	return user, r.Update(c, user)
}

func (r userRepository) CreateFirebaseUser(c context.Context, firebaseID domain.FirebaseID) (*domain.User, error) {
	user := domain.NewFirebaseUser(firebaseID)
	return user, r.Update(c, user)
}

func (r userRepository) Update(c context.Context, user *domain.User) error {
	key := newKey(userKind, string(user.ID), nil)
	return runWithClient(c, func(cli *datastore.Client) error {
		_, err := cli.Put(c, key, user)
		return err
	})
}

func (r userRepository) delete(c context.Context, userID domain.UserID) error {
	key := newKey(userKind, string(userID), nil)
	return runWithClient(c, func(cli *datastore.Client) error {
		return cli.Delete(c, key)
	})
}
