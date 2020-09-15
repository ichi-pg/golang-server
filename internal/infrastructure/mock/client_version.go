package mock

import (
	"github.com/ichi-pg/golang-middleware/repository"
)

type clientVersionRepository struct {
}

// ClientVersionRepository はクライアントバージョンの取得を実装します。
func ClientVersionRepository() repository.ClientVersionRepository {
	return clientVersionRepository{}
}

func (clientVersionRepository) Version(platform string) int {
	return 0
}
