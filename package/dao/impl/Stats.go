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
	cache *cache.TennisMomentCacheDAOImpl
}

func NewStatsDaoImpl() *StatsDaoImpl {
	return &StatsDaoImpl{db: config.DB, cache: cache.NewTennisMomentCacheDAOImpl()}
}

func (impl *StatsDaoImpl) CreateStatsInfo(ctx context.Context, Stats model.Stats) bool {
	impl.db.Save(&Stats)
	return true
}

func (impl *StatsDaoImpl) GetStatsInfoByUid(ctx context.Context, uId int64) model.Stats {
	var Stats model.Stats
	impl.db.First(&Stats, "id", uId)
	return Stats
}

func (impl *StatsDaoImpl) FindAllStats(ctx context.Context, page int, limit int) []model.Stats {
	Statss := make([]model.Stats, 0)
	if page <= 0 || limit <= 0 {
		impl.db.Find(&Statss)
	} else {
		impl.db.Limit(limit).Offset((page - 1) * limit).Find(&Statss)
	}
	return Statss
}

func (impl *StatsDaoImpl) UpdateStatsInfoById(ctx context.Context, Stats model.Stats) bool {
	impl.db.Model(&model.Stats{}).Where("id = ?", Stats.PlayerId).Updates(Stats)
	return true
}

// func (impl *StatsDaoImpl) GetStatsByLoginName(ctx context.Context, loginName string) model.Stats {
// 	var Stats model.Stats
// 	impl.db.First(&Stats, "login_name", loginName)
// 	return Stats
// }
