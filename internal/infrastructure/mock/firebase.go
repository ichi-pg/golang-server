package mock

import (
	"context"
	"net/http"

	"github.com/ichi-pg/golang-server/internal/domain"
)

// FirebaseID はモックにおける正常系のIDです。
const FirebaseID = domain.FirebaseID("74922948-a52b-4f7e-8d25-ffd02c9d0c44")

// FirebaseToken はモックにおける正常系のトークンです。
const FirebaseToken = domain.FirebaseToken("a7fa6b3b-c854-4a7b-b778-152cf55caf60")

// NewFirebaseID はモックにおける正常系の新規のIDです。
const NewFirebaseID = domain.FirebaseID("c49ea001-9dff-4b42-9d8a-9dee4348d0d5")

// NewFirebaseToken はモックにおける正常系の新規のトークンです。
const NewFirebaseToken = domain.FirebaseToken("ffbe0e35-46b3-4a07-bc6e-6837c1cb12d8")

type firebaseService struct {
}

// FirebaseRepository はFirebaseRepositoryのモック実装を返します。
func FirebaseRepository() domain.FirebaseRepository {
	return firebaseService{}
}

func (s firebaseService) FirebaseID(c context.Context, token domain.FirebaseToken) (domain.FirebaseID, error) {
	if token == FirebaseToken {
		return FirebaseID, nil
	}
	if token == NewFirebaseToken {
		return NewFirebaseID, nil
	}
	return domain.FirebaseID(""), domain.NewRequestError(http.StatusUnauthorized)
}
