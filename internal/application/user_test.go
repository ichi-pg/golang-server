package application

import (
	"context"
	"testing"

	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/infrastructure/mock"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_UserAuth(t *testing.T) {
	t.Parallel()

	c := NewAuthContext(context.Background(), logrus.NewEntry(logrus.New()))
	u := UserUsecase{
		repo: mock.UserRepository(),
		fb:   mock.FirebaseRepository(),
	}

	// Test case: Valid.
	user, err := u.UserAuth(c, mock.User.ID, mock.User.Token)
	if assert.NoError(t, err) {
		assert.Equal(t, user.ID, mock.User.ID)
		assert.Equal(t, user.Token, mock.User.Token)
	}

	// Test case: Invalid userID.
	userID := domain.UserID(uuid.NewV4().String())
	_, err = u.UserAuth(c, userID, mock.User.Token)
	assert.Error(t, err)

	// Test case: Invalid token.
	token := domain.UserToken(uuid.NewV4().String())
	_, err = u.UserAuth(c, mock.User.ID, token)
	assert.Error(t, err)
}

func TestUserUsecase_DummyAuth(t *testing.T) {
	t.Parallel()

	c := NewAuthContext(context.Background(), logrus.NewEntry(logrus.New()))
	u := UserUsecase{
		repo: mock.UserRepository(),
		fb:   mock.FirebaseRepository(),
	}

	// Test case: Exists ID.
	user, err := u.DummyAuth(c, mock.User.DummyID)
	if assert.NoError(t, err) {
		assert.Equal(t, user.ID, mock.User.ID)
		assert.Equal(t, user.DummyID, mock.User.DummyID)
	}

	// Test case: New ID.
	dummyID := domain.DummyID(uuid.NewV4().String())
	user, err = u.DummyAuth(c, dummyID)
	if assert.NoError(t, err) {
		assert.NotEqual(t, user.ID, mock.User.ID)
		assert.NotEqual(t, user.DummyID, mock.User.DummyID)
		assert.Equal(t, user.DummyID, dummyID)
	}

	// Test case: Invalid.
	_, err = u.DummyAuth(c, "")
	assert.Error(t, err)
}

func TestUserUsecase_FirebaseAuth(t *testing.T) {
	t.Parallel()

	c := NewAuthContext(context.Background(), logrus.NewEntry(logrus.New()))
	u := UserUsecase{
		repo: mock.UserRepository(),
		fb:   mock.FirebaseRepository(),
	}

	// Test case: Exists ID.
	user, err := u.FirebaseAuth(c, mock.FirebaseToken)
	if assert.NoError(t, err) {
		assert.Equal(t, user.ID, mock.User.ID)
		assert.Equal(t, user.FirebaseID, mock.User.FirebaseID)
	}

	// Test case: New ID.
	user, err = u.FirebaseAuth(c, mock.NewFirebaseToken)
	if assert.NoError(t, err) {
		assert.NotEqual(t, user.ID, mock.User.ID)
		assert.NotEqual(t, user.FirebaseID, mock.User.FirebaseID)
		assert.Equal(t, user.FirebaseID, mock.NewFirebaseID)
	}

	// Test case: Invalid.
	firebaseToken := domain.FirebaseToken(uuid.NewV4().String())
	_, err = u.FirebaseAuth(c, firebaseToken)
	assert.Error(t, err)
}
