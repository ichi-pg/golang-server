package firebases

import (
	"context"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/pkg/env"
	"google.golang.org/api/option"
)

type firebaseService struct {
}

// FirebaseRepository はFirebaseRepositoryの実装を返します。
func FirebaseRepository() domain.FirebaseRepository {
	return firebaseService{}
}

func (s firebaseService) FirebaseID(c context.Context, token domain.FirebaseToken) (domain.FirebaseID, error) {
	opt := option.WithCredentialsJSON([]byte(os.Getenv(env.FirebaseJSON)))
	app, err := firebase.NewApp(c, nil, opt)
	if err != nil {
		return domain.FirebaseID(""), err
	}
	auth, err := app.Auth(c)
	if err != nil {
		return domain.FirebaseID(""), err
	}
	res, err := auth.VerifyIDToken(c, string(token))
	if err != nil {
		return domain.FirebaseID(""), domain.NewRequestError(http.StatusUnauthorized, err.Error())
	}
	return domain.FirebaseID(res.UID), nil
}
