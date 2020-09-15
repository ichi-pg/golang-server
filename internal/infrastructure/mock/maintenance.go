package mock

import (
	"github.com/ichi-pg/golang-middleware/repository"
)

type maintenanceRepository struct {
}

// MaintenanceRepository はクライアントバージョンの取得を実装します。
func MaintenanceRepository() repository.MaintenanceRepository {
	return maintenanceRepository{}
}

func (maintenanceRepository) Active() bool {
	return false
}

func (maintenanceRepository) Message() string {
	return ""
}
