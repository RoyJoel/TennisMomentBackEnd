package impl

import (
	"context"

	"github.com/RoyJoel/TennisMomentBackEnd/package/cache"
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"gorm.io/gorm"
)

type PlayerDaoImpl struct {
	db    *gorm.DB
	cache *cache.TennisMomentCacheDAOImpl
}

func NewPlayerDaoImpl() *PlayerDaoImpl {
	return &PlayerDaoImpl{db: config.DB, cache: cache.NewTennisMomentCacheDAOImpl()}
}

func (impl *PlayerDaoImpl) CreatePlayer(ctx context.Context, Player model.Player) bool {
	var p model.Player
	impl.db.First(&p, "login_name", Player.LoginName)
	if p.LoginName == Player.LoginName {
		return false
	}

	impl.db.Save(&Player)
	return true
}

func (impl *PlayerDaoImpl) GetPlayerInfoByUid(ctx context.Context, uId int64) model.Player {
	var Player model.Player
	impl.db.First(&Player, "id", uId)
	return Player
}

func (impl *PlayerDaoImpl) GetAll(ctx context.Context, page int, limit int) []model.Player {
	Players := make([]model.Player, 0)
	if page <= 0 || limit <= 0 {
		impl.db.Find(&Players)
	} else {
		impl.db.Limit(limit).Offset((page - 1) * limit).Find(&Players)
	}
	return Players
}

func (impl *PlayerDaoImpl) UpdatePlayerById(ctx context.Context, Player model.Player) bool {
	impl.db.Model(&model.Player{}).Where("id = ?", Player.Id).Updates(Player)
	return true
}

func (impl *PlayerDaoImpl) GetPlayerByLoginName(ctx context.Context, loginName string) model.Player {
	var Player model.Player
	impl.db.First(&Player, "loginName", loginName)
	return Player
}
