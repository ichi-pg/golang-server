package mock

import (
	"context"
	"time"

	"github.com/ichi-pg/golang-server/internal/domain"
)

// UserID はモックにおける正常系のIDです。
const UserID = domain.UserID("97c8c495-9ee4-4e23-a97c-9c46ec98aa98")

// UserToken はモックにおける正常系のトークンです。
const UserToken = domain.UserToken("1e07439b-84b6-477c-97f4-f9ccdebede61")

// DummyID はモックにおける正常系のIDです。
const DummyID = domain.DummyID("0ebe6de4-5642-442f-b4f7-ee368614342e")

// UserName はモックにおける正常系のユーザー名です。
const UserName = domain.UserName("hoge")

// UserCreateAt はモックにおける正常系のユーザー作成日です。
var UserCreateAt = time.Date(2020, 5, 1, 9, 0, 0, 0, time.UTC)

// NewUser はモックにおける正常系のユーザーを返します。
func NewUser() *domain.User {
	return &domain.User{
		ID:         UserID,
		Token:      UserToken,
		FirebaseID: FirebaseID,
		DummyID:    DummyID,
		Name:       UserName,
		CreatedAt:  UserCreateAt,
	}
}

type userRepository struct {
}

// UserRepository はUserRepositoryのモック実装を返します。
func UserRepository() domain.UserRepository {
	return userRepository{}
}

func (r userRepository) ByUserID(c context.Context, userID domain.UserID) (*domain.User, error) {
	if userID != UserID {
		return nil, domain.ErrNoSuchEntity
	}
	return NewUser(), nil
}

func (r userRepository) ByDummyID(c context.Context, dummyID domain.DummyID) (*domain.User, error) {
	if dummyID != DummyID {
		return nil, domain.ErrNoSuchEntity
	}
	return NewUser(), nil
}

func (r userRepository) ByFirebaseID(c context.Context, firebaseID domain.FirebaseID) (*domain.User, error) {
	if firebaseID != FirebaseID {
		return nil, domain.ErrNoSuchEntity
	}
	return NewUser(), nil
}

func (r userRepository) CreateDummyUser(c context.Context, dummyID domain.DummyID) (*domain.User, error) {
	return domain.NewDummyUser(dummyID), nil
}

func (r userRepository) CreateFirebaseUser(c context.Context, firebaseID domain.FirebaseID) (*domain.User, error) {
	return domain.NewFirebaseUser(firebaseID), nil
}

func (r userRepository) Update(c context.Context, user *domain.User) error {
	return nil
}
