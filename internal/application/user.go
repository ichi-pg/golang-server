package application

import (
	"net/http"

	"github.com/ichi-pg/golang-server/internal/domain"
)

// UserUsecase はユーザーのユースケースを実装します。
type UserUsecase struct {
	repo domain.UserRepository
	fb   domain.FirebaseRepository
}

// NewUserUsecase はユーザーのユースケースを生成します。
func NewUserUsecase(repo domain.UserRepository, fb domain.FirebaseRepository) UserUsecase {
	return UserUsecase{
		repo: repo,
		fb:   fb,
	}
}

// UserAuth はユーザートークンを認証します。
func (u UserUsecase) UserAuth(c AuthContext, userID domain.UserID, token domain.UserToken) (*domain.User, error) {
	user, err := u.repo.ByUserID(c.Context, userID)
	if err == domain.ErrNoSuchEntity {
		return nil, domain.NewRequestError(http.StatusUnauthorized)
	}
	if err != nil {
		return nil, err
	}
	return user, user.Auth(token)
}

// DummyAuth はダミーユーザーを認証します。
func (u UserUsecase) DummyAuth(c AuthContext, dummyID domain.DummyID) (*domain.User, error) {
	if err := dummyID.Check(); err != nil {
		return nil, err
	}
	user, err := u.repo.ByDummyID(c.Context, dummyID)
	if err == domain.ErrNoSuchEntity {
		user, err = u.repo.CreateDummyUser(c.Context, dummyID)
	}
	return user, err
}

// FirebaseAuth はFirebaseトークンを認証します。
func (u UserUsecase) FirebaseAuth(c AuthContext, token domain.FirebaseToken) (*domain.User, error) {
	firebaseID, err := u.fb.FirebaseID(c.Context, token)
	if err != nil {
		return nil, err
	}
	user, err := u.repo.ByFirebaseID(c.Context, firebaseID)
	if err == domain.ErrNoSuchEntity {
		user, err = u.repo.CreateFirebaseUser(c.Context, firebaseID)
	}
	return user, err
}

// UpdateName はユーザー名を変更します。
func (u UserUsecase) UpdateName(c UserContext, name domain.UserName) (*domain.User, error) {
	if err := name.Check(); err != nil {
		return nil, err
	}
	if err := c.User.UpdateName(name); err != nil {
		return nil, err
	}
	return c.User, u.repo.Update(c.Context, c.User)
}
