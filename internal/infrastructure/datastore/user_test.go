package datastore

import (
	"context"
	"testing"

	"github.com/ichi-pg/golang-server/internal/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_ByUserID(t *testing.T) {
	t.Parallel()

	c := context.Background()
	r := userRepository{}

	// Test case: Before created.
	userID := domain.UserID(uuid.NewV4().String())
	_, err := r.ByUserID(c, userID)
	assert.Equal(t, err, domain.ErrNoSuchEntity)

	dummyID := domain.DummyID(uuid.NewV4().String())
	user, err := r.CreateDummyUser(c, dummyID)
	assert.NoError(t, err)

	// Test case: After created.
	userID = user.ID
	user, err = r.ByUserID(c, userID)
	if assert.NoError(t, err) {
		assert.Equal(t, user.ID, userID)
	}

	err = r.delete(c, userID)
	assert.NoError(t, err)

	// Test case: After deleted.
	_, err = r.ByUserID(c, userID)
	assert.Equal(t, err, domain.ErrNoSuchEntity)
}

func TestUserRepository_ByDummyID_CreateDummyUser(t *testing.T) {
	t.Parallel()

	c := context.Background()
	r := userRepository{}

	// Test case: Before created.
	dummyID := domain.DummyID(uuid.NewV4().String())
	user, err := r.ByDummyID(c, dummyID)
	assert.Equal(t, err, domain.ErrNoSuchEntity)

	// Test case: Create.
	user, err = r.CreateDummyUser(c, dummyID)
	if assert.NoError(t, err) {
		assert.Equal(t, user.DummyID, dummyID)
	}

	// Test case: After created.
	user, err = r.ByDummyID(c, dummyID)
	if assert.NoError(t, err) {
		assert.Equal(t, user.DummyID, dummyID)
	}

	err = r.delete(c, user.ID)
	assert.NoError(t, err)

	// Test case: After deleted.
	_, err = r.ByDummyID(c, dummyID)
	assert.Equal(t, err, domain.ErrNoSuchEntity)
}

func TestUserRepository_ByFirebaseID_CreateFirebaseUser(t *testing.T) {
	t.Parallel()

	c := context.Background()
	r := userRepository{}

	// Test case: Before created.
	firebaseID := domain.FirebaseID(uuid.NewV4().String())
	user, err := r.ByFirebaseID(c, firebaseID)
	assert.Equal(t, err, domain.ErrNoSuchEntity)

	// Test case: Create.
	user, err = r.CreateFirebaseUser(c, firebaseID)
	if assert.NoError(t, err) {
		assert.Equal(t, user.FirebaseID, firebaseID)
	}

	// Test case: After created.
	user, err = r.ByFirebaseID(c, firebaseID)
	if assert.NoError(t, err) {
		assert.Equal(t, user.FirebaseID, firebaseID)
	}

	err = r.delete(c, user.ID)
	assert.NoError(t, err)

	// Test case: After deleted.
	_, err = r.ByFirebaseID(c, firebaseID)
	assert.Equal(t, err, domain.ErrNoSuchEntity)
}
