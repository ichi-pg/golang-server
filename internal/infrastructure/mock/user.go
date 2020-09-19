package mock

import (
	"context"
	"time"

	"github.com/ichi-pg/golang-server/internal/domain"
)

// User はモックにおける正常系のユーザーです。
var User = domain.User{
	ID:         domain.UserID("97c8c495-9ee4-4e23-a97c-9c46ec98aa98"),
	Token:      domain.UserToken("1e07439b-84b6-477c-97f4-f9ccdebede61"),
	FirebaseID: domain.FirebaseID("74922948-a52b-4f7e-8d25-ffd02c9d0c44"),
	DummyID:    domain.DummyID("0ebe6de4-5642-442f-b4f7-ee368614342e"),
	Name:       domain.UserName("hoge"),
	CreatedAt:  time.Date(2020, 5, 1, 9, 0, 0, 0, time.UTC),
}

type userRepository struct {
}

// UserRepository はUserRepositoryのモック実装を返します。
func UserRepository() domain.UserRepository {
	return userRepository{}
}

func (r userRepository) ByUserID(c context.Context, userID domain.UserID) (*domain.User, error) {
	if userID == User.ID {
		return &User, nil
	}
	if userID == RankingUserA.ID {
		return &RankingUserA, nil
	}
	if userID == RankingUserB.ID {
		return &RankingUserB, nil
	}
	if userID == RankingUserC.ID {
		return &RankingUserC, nil
	}
	return nil, domain.ErrNoSuchEntity
}

func (r userRepository) ByDummyID(c context.Context, dummyID domain.DummyID) (*domain.User, error) {
	if dummyID != User.DummyID {
		return nil, domain.ErrNoSuchEntity
	}
	return &User, nil
}

func (r userRepository) ByFirebaseID(c context.Context, firebaseID domain.FirebaseID) (*domain.User, error) {
	if firebaseID != User.FirebaseID {
		return nil, domain.ErrNoSuchEntity
	}
	return &User, nil
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
