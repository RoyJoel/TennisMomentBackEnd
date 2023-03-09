package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/RoyJoel/TennisMomentBackEnd/package/cache"
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"gorm.io/gorm"
)

type PlayerStatsDaoImpl struct {
	db    *gorm.DB
	cache *cache.PlayerStatsCacheDAOImpl
}

func NewPlayerStatsDaoImpl() *PlayerStatsDaoImpl {
	return &PlayerStatsDaoImpl{db: config.DB, cache: cache.NewPlayerStatsCacheDAOImpl()}
}

func (impl *PlayerStatsDaoImpl) AddPlayerStats(ctx context.Context, PlayerStats model.PlayerStats) bool {
	impl.db.Save(&PlayerStats)
	return true
}

// func (impl *PlayerStatsDaoImpl) AddPlayerStats(ctx context.Context, PlayerStats model.PlayerStats) bool {
// 	var p model.PlayerStats
// 	impl.db.First(&p, "loginName", PlayerStats.LoginName)
// 	if p.LoginName == PlayerStats.LoginName {
// 		return false
// 	}

// 	impl.db.Save(&PlayerStats)
// 	return true
// }

func (impl *PlayerStatsDaoImpl) UpdatePlayerStats(ctx context.Context, PlayerStats model.PlayerStats) bool {
	result := impl.db.Model(&model.PlayerStats{}).Where("login_name = ? ", PlayerStats.LoginName).Updates(PlayerStats)
	return result.RowsAffected > 0
}

func (impl *PlayerStatsDaoImpl) SearchPlayerStats(ctx context.Context, name string) (*model.PlayerStats, error) {
	var PlayerStats model.PlayerStats
	result := impl.db.First(&PlayerStats, "login_name = ", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("PlayerStats not found for players: %s", name)
		}
		return nil, fmt.Errorf("failed to get PlayerStats for players %s: %v", name, result.Error)
	}
	return &PlayerStats, nil
}
