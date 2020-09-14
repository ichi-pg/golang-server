package domain

import (
	"context"
)

// FirebaseRepository はFirebaseの操作を抽象化します。
type FirebaseRepository interface {
	FirebaseID(c context.Context, token FirebaseToken) (FirebaseID, error)
}
