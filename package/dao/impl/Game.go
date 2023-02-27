// package impl

// import (
// 	"TennisMoment/package/cache"
// 	"TennisMoment/package/config"
// 	"TennisMoment/package/model"
// 	"TennisMoment/package/utils"
// 	"context"

// 	"gorm.io/gorm"
// )

// type GameDaoImpl struct {
// 	db    *gorm.DB
// 	cache *cache.TennisMomentCacheDAOImpl
// }

// func NewGameDaoImpl() *GameDaoImpl {
// 	return &GameDaoImpl{db: config.DB, cache: cache.NewTennisMomentCacheDAOImpl()}
// }

// func (impl *GameDaoImpl) CreateGame(ctx context.Context, Game model.Game) bool {
// 	Game.date = utils.NowTimeStr
// 	Game.place = Game.place
// 	Game.winner = Game.winner
// 	Game.winnerStats = Game.winnerStats
// 	Game.loser = Game.loser
// 	Game.loserStats = Game.loserStats

// 	impl.db.Save(&Game)
// 	return true
// }

// func (impl *GameDaoImpl) GetGameByPid(ctx context.Context, uId int64) model.Game {
// 	var Game model.Game
// 	impl.db.First(&Game, "id", uId)
// 	return Game
// }

// func (impl *GameDaoImpl) GetAll(ctx context.Context, page int, limit int) []model.Game {
// 	Games := make([]model.Game, 0)
// 	if page <= 0 || limit <= 0 {
// 		impl.db.Find(&Games)
// 	} else {
// 		impl.db.Limit(limit).Offset((page - 1) * limit).Find(&Games)
// 	}
// 	return Games
// }

// func (impl *GameDaoImpl) UpdateGameById(ctx context.Context, Game model.Game) bool {
// 	impl.db.Model(&model.Game{}).Where("id = ?", Game.id).Updates(Game)
// 	return true
// }

package impl