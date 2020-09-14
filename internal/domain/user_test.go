package domain

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestDummyID_Check(t *testing.T) {
	t.Parallel()

	// Test case: Valid.
	assert.NoError(t, DummyID(uuid.NewV4().String()).Check())

	// Test case: Invalid.
	assert.Error(t, DummyID("").Check())
}

func TestUser_Auth(t *testing.T) {
	t.Parallel()

	token := UserToken(uuid.NewV4().String())

	user := &User{
		Token: token,
	}

	// Test case: Valid.
	assert.NoError(t, user.Auth(token))

	// Test case: Invalid.
	assert.Error(t, user.Auth((UserToken(uuid.NewV4().String()))))
}

func TestUser_UpdateName(t *testing.T) {
	t.Parallel()

	//TODO
}

func TestNewDummyUser(t *testing.T) {
	t.Parallel()

	start := time.Now()

	dummyID := DummyID(uuid.NewV4().String())
	user := NewDummyUser(dummyID)

	if assert.NotNil(t, user) {
		assert.NotEqual(t, user.ID, UserID(""))
		assert.NotEqual(t, user.Token, UserToken(""))
		assert.Equal(t, user.DummyID, dummyID)
		assert.Equal(t, user.FirebaseID, FirebaseID(""))
		assert.Greater(t, user.CreatedAt.UnixNano(), start.UnixNano())
		assert.Less(t, user.CreatedAt.UnixNano(), time.Now().UnixNano())
	}
}

func TestNewFirebaseUser(t *testing.T) {
	t.Parallel()

	start := time.Now()

	firebaseID := FirebaseID(uuid.NewV4().String())
	user := NewFirebaseUser(firebaseID)

	if assert.NotNil(t, user) {
		assert.NotEqual(t, user.ID, UserID(""))
		assert.NotEqual(t, user.Token, UserToken(""))
		assert.Equal(t, user.DummyID, DummyID(""))
		assert.Equal(t, user.FirebaseID, firebaseID)
		assert.Greater(t, user.CreatedAt.UnixNano(), start.UnixNano())
		assert.Less(t, user.CreatedAt.UnixNano(), time.Now().UnixNano())
	}
}
