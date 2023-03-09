package impl

import (
	"context"

	"github.com/RoyJoel/TennisMomentBackEnd/package/cache"
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"

	"gorm.io/gorm"
)

type GameDaoImpl struct {
	db    *gorm.DB
	cache *cache.GameCacheDAOImpl
}

func NewGameDaoImpl() *GameDaoImpl {
	return &GameDaoImpl{db: config.DB, cache: cache.NewGameCacheDAOImpl()}
}

func (impl *GameDaoImpl) AddGame(ctx context.Context, Game *model.Game) bool {
	Stats1 := model.Stats{}
	Stats2 := model.Stats{}

	impl.db.Create(&Stats1)
	impl.db.Create(&Stats2)

	PlayerStats1 := model.PlayerStats{LoginName: Game.Player1LoginName, Id: int(Stats1.Id)}
	PlayerStats2 := model.PlayerStats{LoginName: Game.Player2LoginName, Id: int(Stats2.Id)}

	impl.db.Create(&PlayerStats1)
	impl.db.Create(&PlayerStats2)

	Game.Player1StatsId = int(Stats1.Id)
	Game.Player2StatsId = int(Stats2.Id)

	impl.db.Create(&Game)
	return true
}

// func (impl *GameDaoImpl) AddGame(ctx context.Context, Game model.Game) bool {
// 	var p model.Game
// 	impl.db.First(&p, "loginName", Game.LoginName)
// 	if p.LoginName == Game.LoginName {
// 		return false
// 	}

// 	impl.db.Save(&Game)
// 	return true
// }

func (impl *GameDaoImpl) UpdateGame(ctx context.Context, date float64, place string, surface string, isCompleted bool, player1LoginName string, player1StatsId int, player2LoginName string, result utils.IntMatrix) model.Game {
	Game := model.Game{}
	impl.db.First(&Game, "player1_login_name = ? AND player2_login_name = ? AND player1_stats_id = ?", player1LoginName, player2LoginName, player1StatsId)
	Game.Date = date
	Game.Place = place
	Game.Surface = surface
	Game.IsCompleted = isCompleted
	Game.Result = result
	impl.db.Where("player1_login_name = ? AND player2_login_name = ? AND player1_stats_id = ?", player1LoginName, player2LoginName, player1StatsId).Save(Game)
	return Game
}

func (impl *GameDaoImpl) SearchGame(ctx context.Context, player1LoginName, player2LoginName string) []model.Game {
	var game1 []model.Game
	var game2 []model.Game
	impl.db.Find(&game1, "player1_login_name = ? AND player2_login_name = ?", player1LoginName, player2LoginName)
	impl.db.Find(&game2, "player1_login_name = ? AND player2_login_name = ?", player2LoginName, player1LoginName)

	for _, game := range game2 {
		game1 = append(game1, game)
	}

	return game1
}

func (impl *GameDaoImpl) GetAllGames(ctx context.Context, loginName string) []model.Game {
	var game1 []model.Game
	var game2 []model.Game
	impl.db.Find(&game1, "player1_login_name = ?", loginName)
	impl.db.Find(&game2, "player2_login_name = ?", loginName)
	for _, game := range game2 {
		game1 = append(game1, game)
	}

	return game1
}
