package impl

import (
	"context"

	"github.com/RoyJoel/TennisMomentBackEnd/package/cache"
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"gorm.io/gorm"
)

type StatsDaoImpl struct {
	db    *gorm.DB
	cache *cache.StatsCacheDAOImpl
}

func NewStatsDaoImpl() *StatsDaoImpl {
	return &StatsDaoImpl{db: config.DB, cache: cache.NewStatsCacheDAOImpl()}
}

// func (impl *StatsDaoImpl) CreateStatsByPlayerLoginName(ctx context.Context, playerLoginName string) bool {
// 	stats := model.Stats{}
// 	stats.PlayerLoginName = playerLoginName
// 	impl.db.Create(stats)
// 	return true
// }

func (impl *StatsDaoImpl) SearchStatsInfo(ctx context.Context, uId int64) model.Stats {
	var Stats model.Stats
	impl.db.First(&Stats, "id", uId)
	return Stats
}

func (impl *StatsDaoImpl) GetStatsInfoByPlayerLoginName(ctx context.Context, playerLoginName string) model.Stats {
	var Stats model.Stats
	impl.db.First(&Stats, "id", playerLoginName)
	return Stats
}

// func (impl *StatsDaoImpl) searchStats(ctx context.Context, page int, limit int) model.Stats {
// 	Stats := model.Stats{}
// 	if page <= 0 || limit <= 0 {
// 		impl.db.Find(&Stats)
// 	} else {
// 		impl.db.Limit(limit).Offset((page - 1) * limit).Find(&Statss)
// 	}
// 	return Statss
// }

func (impl *StatsDaoImpl) UpdateStatsInfoById(ctx context.Context, Stats model.Stats) bool {
	impl.db.Save(Stats)
	return true
}

// func (impl *StatsDaoImpl) GetStatsByLoginName(ctx context.Context, loginName string) model.Stats {
// 	var Stats model.Stats
// 	impl.db.First(&Stats, "login_name", loginName)
// 	return Stats
// }
