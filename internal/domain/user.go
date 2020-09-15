package domain

import (
	"context"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	// UserID はユーザーの永続化IDです。
	UserID string

	// UserToken はユーザー認証トークンです。
	UserToken string

	// FirebaseID はFirebaseの認証IDです。
	FirebaseID string

	// FirebaseToken はFirebaseトークンです。
	FirebaseToken string

	// DummyID はダミーユーザーの認証IDです。
	DummyID string

	// UserName はユーザー名です。
	UserName string
)

// User はユーザーIDによってユニークなユーザーデータです。
type User struct {
	ID         UserID
	Token      UserToken
	FirebaseID FirebaseID
	DummyID    DummyID
	Name       UserName
	CreatedAt  time.Time
}

// UserRepository はユーザーのCRUDを抽象化します。
type UserRepository interface {
	ByUserID(c context.Context, userID UserID) (*User, error)
	ByDummyID(c context.Context, dummyID DummyID) (*User, error)
	ByFirebaseID(c context.Context, firebaseID FirebaseID) (*User, error)
	CreateDummyUser(c context.Context, dummyID DummyID) (*User, error)
	CreateFirebaseUser(fc context.Context, irebaseID FirebaseID) (*User, error)
	Update(c context.Context, user *User) error
}

// Check はダミーユーザーIDをチェックします。
func (id DummyID) Check() error {
	if id == "" {
		return NewRequestError(http.StatusBadRequest, "ダミーIDが空です。")
	}
	return nil
}

// Check はユーザー名をチェックします。
func (name UserName) Check() error {
	if name == "" {
		return NewRequestError(http.StatusBadRequest, "ユーザー名が空です。")
	}
	return nil
}

// Auth はユーザートークンをチェックします。
func (user *User) Auth(token UserToken) error {
	if user.Token != token {
		return NewRequestError(http.StatusUnauthorized)
	}
	return nil
}

// UpdateName はユーザー名を変更します。
func (user *User) UpdateName(name UserName) error {
	if err := name.Check(); err != nil {
		return err
	}
	user.Name = name
	return nil
}

// NewDummyUser はダミーユーザーを作成します。
func NewDummyUser(dummyID DummyID) *User {
	user := newUser()
	user.DummyID = dummyID
	return user
}

// NewFirebaseUser はFirebaseユーザーを作成します。
func NewFirebaseUser(firebaseID FirebaseID) *User {
	user := newUser()
	user.FirebaseID = firebaseID
	return user
}

func newUser() *User {
	s := uuid.NewV4().String()
	return &User{
		ID:        UserID(s),
		Name:      UserName(s),
		Token:     UserToken(uuid.NewV4().String()),
		CreatedAt: time.Now(),
	}
}
