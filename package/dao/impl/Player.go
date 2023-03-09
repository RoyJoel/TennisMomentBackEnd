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
	cache *cache.PlayerCacheDAOImpl
}

func NewPlayerDaoImpl() *PlayerDaoImpl {
	return &PlayerDaoImpl{db: config.DB, cache: cache.NewPlayerCacheDAOImpl()}
}

func (impl *PlayerDaoImpl) AddPlayer(ctx context.Context, Player model.Player) model.Player {

	res := impl.SearchPlayer(ctx, Player.LoginName)
	if !res {
		Stats := model.Stats{}

		impl.db.Create(&Player)
		impl.db.Create(&Stats)
		PlayerStats := model.PlayerStats{LoginName: Player.LoginName, Id: int(Stats.Id)}
		impl.db.Create(&PlayerStats)
		Player.CareerStatsId = int(Stats.Id)
		return Player
	}
	return model.Player{}
}

func (impl *PlayerDaoImpl) UpdatePlayer(ctx context.Context, Player model.Player) model.Player {
	impl.db.Model(&model.Player{}).Where("login_name = ?", Player.LoginName).Updates(Player)
	impl.db.First(&Player, "login_name", Player.LoginName)
	return Player
}

func (impl *PlayerDaoImpl) SearchPlayer(ctx context.Context, loginName string) bool {
	var Player model.Player
	impl.db.Where("login_name = ?", loginName).First(&Player)
	return Player.LoginName == loginName
}

func (impl *PlayerDaoImpl) GetPlayerInfo(ctx context.Context, loginName string) model.Player {
	var Player model.Player
	impl.db.Where("login_name = ?", loginName).First(&Player)
	if Player.LoginName == loginName {
		return Player
	}
	return model.Player{}
}

func (impl *PlayerDaoImpl) GetAllFriends(ctx context.Context, loginName string) []model.Relationship {
	var relationship []model.Relationship
	impl.db.Find(&relationship, "login_name", loginName)
	return relationship
}

func (impl *PlayerDaoImpl) AddFriend(ctx context.Context, relationship model.Relationship) []model.Relationship {
	res := impl.SearchFriend(ctx, relationship)
	if !res {
		impl.db.Create(&relationship)
	}
	return impl.GetAllFriends(ctx, relationship.LoginName)
}

func (impl *PlayerDaoImpl) DeleteFriend(ctx context.Context, relationship model.Relationship) []model.Relationship {
	if impl.SearchFriend(ctx, relationship) {
		impl.db.Where("login_name = ? AND friend_login_name = ?", relationship.LoginName, relationship.FriendLoginName).Delete(&relationship)
	}
	return impl.GetAllFriends(ctx, relationship.LoginName)
}

func (impl *PlayerDaoImpl) SearchFriend(ctx context.Context, relationship model.Relationship) bool {
	var Relationship model.Relationship
	impl.db.Where("login_name = ? AND friend_login_name = ?", relationship.LoginName, relationship.FriendLoginName).First(&Relationship)
	return Relationship == relationship

}

// func (impl *PlayerDaoImpl) FindAllPlayers(ctx context.Context, page int, limit int) []model.Player {
// 	var players []model.Player
// 	if page <= 0 || limit <= 0 {
// 		impl.db.Find(&players)
// 	} else {
// 		impl.db.Limit(limit).Offset((page - 1) * limit).Find(&players)
// 	}
// 	return players
// }

// func (impl *PlayerDaoImpl) GetPlayerInfoByUid(ctx context.Context, uId int64) model.Player {
// 	var Player model.Player
// 	impl.db.First(&Player, "id", uId)
// 	return Player
// }

// func (impl *PlayerDaoImpl) GetAll(ctx context.Context, page int, limit int) []model.Player {
// 	Players := make([]model.Player, 0)
// 	if page <= 0 || limit <= 0 {
// 		impl.db.Find(&Players)
// 	} else {
// 		impl.db.Limit(limit).Offset((page - 1) * limit).Find(&Players)
// 	}
// 	return Players
// }
