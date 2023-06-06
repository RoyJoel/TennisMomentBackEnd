package impl

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/RoyJoel/TennisMomentBackEnd/package/cache"
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/middleware"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"
	"gorm.io/gorm"
)

type TennisMomentDaoImpl struct {
	db    *gorm.DB
	cache *cache.TennisMomentCacheDAOImpl
}

func NewTennisMomentDaoImpl() *TennisMomentDaoImpl {
	return &TennisMomentDaoImpl{db: config.DB, cache: cache.NewTennisMomentCacheDAOImpl()}
}

func (impl *TennisMomentDaoImpl) AddGame(ctx context.Context, game *model.Game) model.GameResponse {
	gameResponse := model.GameResponse{}

	Stats1 := model.Stats{}
	Stats2 := model.Stats{}

	player1, _ := impl.GetPlayerInfo(ctx, game.Player1Id)
	player2, _ := impl.GetPlayerInfo(ctx, game.Player2Id)

	impl.db.Create(&Stats1)
	impl.db.Create(&Stats2)

	PlayerStats1 := model.PlayerStats{PlayerId: game.Player1Id, StatsId: Stats1.Id}
	PlayerStats2 := model.PlayerStats{PlayerId: game.Player2Id, StatsId: Stats2.Id}

	impl.db.Create(&PlayerStats1)
	impl.db.Create(&PlayerStats2)

	game.Player1StatsId = Stats1.Id
	game.Player2StatsId = Stats2.Id

	impl.db.Create(&game)

	gameResponse = model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: *player1, Player1Stats: Stats1, Player2: *player2, Player2Stats: Stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
	return *&gameResponse
}

func (impl *TennisMomentDaoImpl) assignIdAndSaveGame(ctx context.Context, game *model.GameResponse) {
	newGame := model.Game{}
	newGame.Result = utils.IntMatrix3{{{0, 0}}}
	impl.db.Create(&newGame)

	fmt.Println(newGame)

	Stats1 := model.Stats{}
	Stats2 := model.Stats{}

	impl.db.Create(&Stats1)
	impl.db.Create(&Stats2)

	PlayerStats1 := model.PlayerStats{PlayerId: game.Player1.Id, StatsId: Stats1.Id}
	PlayerStats2 := model.PlayerStats{PlayerId: game.Player2.Id, StatsId: Stats2.Id}

	impl.db.Create(&PlayerStats1)
	impl.db.Create(&PlayerStats2)

	game.Id = newGame.Id
	game.Player1Stats.Id = Stats1.Id
	game.Player2Stats.Id = Stats2.Id

	fmt.Println(game.Id, newGame.Id, game.Player1Stats.Id, game.Player2Stats.Id)
	gameToSave := model.Game{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1Id: game.Player1.Id, Player1StatsId: game.Player1Stats.Id, Player2Id: game.Player2.Id, Player2StatsId: game.Player2Stats.Id, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
	impl.db.Save(&gameToSave)
	impl.db.Save(&game.Player1Stats)
	impl.db.Save(&game.Player2Stats)
	impl.UpdateStats(ctx, game.Player1Stats, game.Player1.CareerStats.Id)
	impl.UpdateStats(ctx, game.Player2Stats, game.Player2.CareerStats.Id)
}

func (impl *TennisMomentDaoImpl) UpdateGameAndStats(ctx context.Context, gameResponse model.GameResponse) model.GameResponse {
	game := model.Game{Id: gameResponse.Id, Place: gameResponse.Place, Surface: gameResponse.Surface, SetNum: gameResponse.SetNum, GameNum: gameResponse.GameNum, Round: gameResponse.Round, IsGoldenGoal: gameResponse.IsGoldenGoal, IsPlayer1Serving: gameResponse.IsPlayer1Serving, IsPlayer1Left: gameResponse.IsPlayer1Left, IsChangePosition: gameResponse.IsChangePosition, StartDate: gameResponse.StartDate, EndDate: gameResponse.EndDate, Player1Id: gameResponse.Player1.Id, Player1StatsId: gameResponse.Player1Stats.Id, Player2Id: gameResponse.Player2.Id, Player2StatsId: gameResponse.Player2Stats.Id, IsPlayer1FirstServe: gameResponse.IsPlayer1FirstServe, IsPlayer2FirstServe: gameResponse.IsPlayer2FirstServe, Result: gameResponse.Result}
	var Game model.Game
	impl.db.Where("id = ?", game.Id).First(&Game)
	if game.Id == 0 {
		impl.assignIdAndSaveGame(ctx, &gameResponse)
	} else if !Game.Equals(game) {
		impl.db.Where("id = ?", game.Id).Save(&game)
		impl.UpdateStats(ctx, gameResponse.Player1Stats, gameResponse.Player1.CareerStats.Id)
		impl.UpdateStats(ctx, gameResponse.Player2Stats, gameResponse.Player2.CareerStats.Id)
	}
	return gameResponse
}

func (impl *TennisMomentDaoImpl) GetRecentGames(ctx context.Context, id int64, limit int, isCompleted bool) []model.GameResponse {
	var games []model.Game
	var gameResponses []model.GameResponse
	impl.db.Order("end_date desc").Limit(limit).Find(&games, "player1_id = ? OR player2_id = ?", id, id)
	for _, game := range games {
		if !isCompleted {
			if game.EndDate >= float64(time.Now().UnixNano()) || game.EndDate == 0 {
				player1, _ := impl.GetPlayerInfo(ctx, game.Player1Id)
				stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
				player2, _ := impl.GetPlayerInfo(ctx, game.Player2Id)
				stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
				gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: *player1, Player1Stats: stats1, Player2: *player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
				gameResponses = append(gameResponses, gameResponse)
			}
		} else {
			if game.EndDate < float64(time.Now().UnixNano()) && game.EndDate != 0 {
				player1, _ := impl.GetPlayerInfo(ctx, game.Player1Id)
				stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
				player2, _ := impl.GetPlayerInfo(ctx, game.Player2Id)
				stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
				gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: *player1, Player1Stats: stats1, Player2: *player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
				gameResponses = append(gameResponses, gameResponse)
			}
		}
	}
	return gameResponses
}

func (impl *TennisMomentDaoImpl) GetHistoryGames(ctx context.Context, player1Id, player2Id int64) []model.GameResponse {
	var game1 []model.Game
	var game2 []model.Game
	var gameResponses []model.GameResponse
	impl.db.Find(&game1, "player1_id = ? AND player2_id = ?", player1Id, player2Id)
	impl.db.Find(&game2, "player2_id = ? AND player1_id = ?", player1Id, player2Id)

	for _, game := range game2 {
		game1 = append(game1, game)
	}

	for _, game := range game1 {
		if game.EndDate < float64(time.Now().UnixNano()) && game.EndDate != 0 {
			player1, _ := impl.GetPlayerInfo(ctx, game.Player1Id)
			stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
			player2, _ := impl.GetPlayerInfo(ctx, game.Player2Id)
			stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
			gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: *player1, Player1Stats: stats1, Player2: *player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
			gameResponses = append(gameResponses, gameResponse)
		}
	}

	return gameResponses
}

func (impl *TennisMomentDaoImpl) GetGameInfo(ctx context.Context, id int64) model.GameResponse {
	var game model.Game
	impl.db.Where("id = ?", id).First(&game)
	if game.Id == id {
		player1, _ := impl.GetPlayerInfo(ctx, game.Player1Id)
		stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
		player2, _ := impl.GetPlayerInfo(ctx, game.Player2Id)
		stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
		return model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: *player1, Player1Stats: stats1, Player2: *player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
	}
	return model.GameResponse{}
}

func (impl *TennisMomentDaoImpl) GetAllGames(ctx context.Context, playerId int64) []model.GameResponse {
	var game1 []model.Game
	var game2 []model.Game
	var gameResponses []model.GameResponse
	impl.db.Find(&game1, "player1_id = ?", playerId)
	impl.db.Find(&game2, "player2_id = ?", playerId)
	for _, game := range game2 {
		game1 = append(game1, game)
	}

	for _, game := range game1 {
		player1, _ := impl.GetPlayerInfo(ctx, game.Player1Id)
		stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
		player2, _ := impl.GetPlayerInfo(ctx, game.Player2Id)
		stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
		gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: *player1, Player1Stats: stats1, Player2: *player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
		gameResponses = append(gameResponses, gameResponse)
	}

	return gameResponses
}

func (impl *TennisMomentDaoImpl) GetAllHistoryGames(ctx context.Context, playerId int64, isCompleted bool) []model.GameResponse {
	var game1 []model.Game
	var game2 []model.Game
	var gameResponses []model.GameResponse
	impl.db.Find(&game1, "player1_id = ?", playerId)
	impl.db.Find(&game2, "player2_id = ?", playerId)
	for _, game := range game2 {
		if game.EndDate == 0 {
			game1 = append(game1, game)
		}
	}

	for _, game := range game1 {
		if !isCompleted {
			if game.EndDate >= float64(time.Now().UnixNano()) || game.EndDate == 0 {
				player1, _ := impl.GetPlayerInfo(ctx, game.Player1Id)
				stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
				player2, _ := impl.GetPlayerInfo(ctx, game.Player2Id)
				stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
				gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: *player1, Player1Stats: stats1, Player2: *player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
				gameResponses = append(gameResponses, gameResponse)
			}
		} else {
			if game.EndDate < float64(time.Now().UnixNano()) && game.EndDate != 0 {
				player1, _ := impl.GetPlayerInfo(ctx, game.Player1Id)
				stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
				player2, _ := impl.GetPlayerInfo(ctx, game.Player2Id)
				stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
				gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: *player1, Player1Stats: stats1, Player2: *player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
				gameResponses = append(gameResponses, gameResponse)
			}
		}
	}

	return gameResponses
}

func (impl *TennisMomentDaoImpl) UpdateGameDB(ctx context.Context, userId int64, games []model.GameResponse, isFinished bool) []model.GameResponse {

	var results []model.Game

	m := make(map[int64]bool)
	impl.db.Where("id = ?", userId, userId).Find(&results)
	for _, relat := range games {
		m[relat.Id] = true
	}
	for _, res := range results {
		if !m[res.Id] {
			impl.db.Where("id = ?", res.Id).Delete(&res)
		}
	}
	return impl.GetAllHistoryGames(ctx, userId, isFinished)
}

func (impl *TennisMomentDaoImpl) AddPlayer(ctx context.Context, Player model.Player) (model.PlayerResponse, bool) {

	res := impl.SearchPlayer(ctx, Player.LoginName)
	if !res {
		newPlayer := model.Player{}
		newPlayer = model.Player{LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, YearsPlayed: Player.YearsPlayed, Height: Player.Height, Width: Player.Width, Grip: Player.Grip, Backhand: Player.Backhand, Points: Player.Points, IsAdult: Player.IsAdult}
		Stats := model.Stats{}
		impl.db.Create(&Stats)
		PlayerStats := model.PlayerStats{PlayerId: newPlayer.Id, StatsId: Stats.Id}
		impl.db.Create(&PlayerStats)
		newPlayer.CareerStatsId = Stats.Id
		impl.db.Create(&newPlayer)
		return model.PlayerResponse{Id: newPlayer.Id, LoginName: newPlayer.LoginName, Password: newPlayer.Password, Name: newPlayer.Name, Icon: newPlayer.Icon, Sex: newPlayer.Sex, Age: newPlayer.Age, YearsPlayed: newPlayer.YearsPlayed, Height: newPlayer.Height, Width: newPlayer.Width, Grip: newPlayer.Grip, Backhand: newPlayer.Backhand, Points: newPlayer.Points, IsAdult: newPlayer.IsAdult, CareerStats: Stats}, res
	} else {
		var player model.Player
		impl.db.Where("id = ?", Player.Id).First(&player)
		stats := impl.GetStatsInfo(ctx, player.CareerStatsId)
		return model.PlayerResponse{Id: player.Id, LoginName: player.LoginName, Password: Player.Password, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, YearsPlayed: player.YearsPlayed, Height: player.Height, Width: player.Width, Grip: player.Grip, Backhand: player.Backhand, Points: player.Points, IsAdult: player.IsAdult, CareerStats: stats}, res
	}
}

func (impl *TennisMomentDaoImpl) SignUp(ctx context.Context, user model.User) (model.User, bool) {
	playerInfo := model.Player{Id: user.Id, LoginName: user.LoginName, Password: user.Password, Name: user.Name, Icon: user.Icon, Sex: user.Sex, Age: user.Age, YearsPlayed: user.YearsPlayed, Height: user.Height, Width: user.Width, Grip: user.Grip, Backhand: user.Backhand, Points: user.Points, IsAdult: user.IsAdult, CareerStatsId: user.CareerStats.Id}
	player, res := impl.AddPlayer(ctx, playerInfo)
	friends := []model.PlayerResponse{}
	if res == false {
		relationship := model.Relationship{Player1Id: player.Id, Player2Id: player.Id}
		impl.db.Create(relationship)
		friends = append(friends, player)
	} else {
		friends = user.Friends
	}
	cart := impl.AddCartForPlayer(ctx, player.Id)
	// 生成 token string
	tokenStr, _ := middleware.GenerateToken(user.LoginName, user.Password)
	user = model.User{Id: player.Id, LoginName: player.LoginName, Password: user.Password, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, YearsPlayed: player.YearsPlayed, Height: player.Height, Width: player.Width, Grip: player.Grip, Backhand: player.Backhand, Points: player.Points, IsAdult: player.IsAdult, CareerStats: player.CareerStats, Friends: friends, AllHistoryGames: make([]model.GameResponse, 0), AllUnfinishedGames: make([]model.GameResponse, 0), AllClubs: utils.IntMatrix{}, AllSchedules: make([]model.ScheduleResponse, 0), AllEvents: utils.IntMatrix{}, AllOrders: utils.IntMatrix{}, Addresss: utils.IntMatrix{}, Cart: cart, DefaultAddress: model.AddressResponse{}, Token: tokenStr}
	return user, res
}
func (impl *TennisMomentDaoImpl) AddCartForPlayer(ctx context.Context, playerId int64) int64 {
	cart := model.Cart{}
	impl.db.Where("player_id = ?", playerId).Delete(&cart)
	newOrder := model.Order{PlayerId: playerId, State: 0}
	impl.db.Create(&newOrder)
	newCart := model.Cart{PlayerId: playerId, OrderId: newOrder.Id}
	impl.db.Create(&newCart)
	return newOrder.Id
}

func (impl *TennisMomentDaoImpl) UpdateUser(ctx context.Context, user model.User) model.User {
	Player := model.PlayerResponse{Id: user.Id, LoginName: user.LoginName, Password: user.Password, Name: user.Name, Icon: user.Icon, Sex: user.Sex, Age: user.Age, YearsPlayed: user.YearsPlayed, Height: user.Height, Width: user.Width, Grip: user.Grip, Backhand: user.Backhand, Points: user.Points, IsAdult: user.IsAdult, CareerStats: user.CareerStats}
	playerResponse := impl.UpdatePlayer(ctx, Player)
	for _, friend := range user.Friends {
		impl.UpdateFriend(ctx, model.Relationship{Player1Id: user.Id, Player2Id: friend.Id})
	}
	user.Friends = impl.UpdateFriendDB(ctx, user.Id, user.Friends)
	for _, club := range user.AllClubs {
		impl.UpdateClub(ctx, model.ClubMember{ClubId: club, MemberId: user.Id})
	}
	user.AllClubs = impl.UpdateClubDB(ctx, user.Id, user.AllClubs)
	for _, game := range user.AllHistoryGames {
		impl.UpdateGameAndStats(ctx, game)
	}
	user.AllHistoryGames = impl.UpdateGameDB(ctx, user.Id, user.AllHistoryGames, true)
	for _, game := range user.AllUnfinishedGames {
		impl.UpdateGameAndStats(ctx, game)
	}
	user.AllUnfinishedGames = impl.UpdateGameDB(ctx, user.Id, user.AllUnfinishedGames, false)
	for _, event := range user.AllEvents {
		impl.UpdateEvent(ctx, model.EventPlayer{EventId: event, PlayerId: user.Id})
	}
	for _, schedule := range user.AllSchedules {
		impl.UpdateSchedule(ctx, model.Schedule{Id: schedule.Id, StartDate: schedule.StartDate, Place: schedule.Place, Player1Id: user.Id, Player2Id: schedule.Opponent.Id})
	}
	user.AllSchedules = impl.UpdateScheduleDB(ctx, user.Id, user.AllSchedules)
	return model.User{Id: playerResponse.Id, LoginName: playerResponse.LoginName, Password: playerResponse.Password, Name: playerResponse.Name, Icon: playerResponse.Icon, Sex: playerResponse.Sex, Age: playerResponse.Age, YearsPlayed: playerResponse.YearsPlayed, Height: playerResponse.Height, Width: playerResponse.Width, Grip: playerResponse.Grip, Backhand: playerResponse.Backhand, Points: playerResponse.Points, IsAdult: playerResponse.IsAdult, CareerStats: playerResponse.CareerStats, Friends: user.Friends, AllClubs: user.AllClubs, AllHistoryGames: user.AllHistoryGames, AllUnfinishedGames: user.AllUnfinishedGames, AllEvents: user.AllEvents, AllSchedules: user.AllSchedules, AllOrders: user.AllOrders, Addresss: user.Addresss, Cart: user.Cart, Token: user.Token}
}

func (impl *TennisMomentDaoImpl) SignIn(ctx context.Context, userLoginName string, userPassword string) (*model.User, error) {
	player, _ := impl.GetPlayerInfoByLoginNameAndPassword(ctx, userLoginName, userPassword)
	if player != nil {
		friends := impl.GetAllFriends(ctx, player.Id)
		games := impl.GetAllHistoryGames(ctx, player.Id, true)
		unfinishedGames := impl.GetAllHistoryGames(ctx, player.Id, false)
		clubs := impl.GetMyClubs(ctx, player.Id)
		schedules := impl.GetSchedules(ctx, player.Id)
		allOrders := impl.GetMyOrders(ctx, player.Id)
		addresss := impl.GetMyAddresss(ctx, player.Id)
		cart := impl.GetMyCart(ctx, player.Id)
		defaultAddress := impl.GetMyDefaultAddress(ctx, player.Id)
		// 生成 token string
		tokenStr, error := middleware.GenerateToken(player.LoginName, player.Password)
		if error != nil {
			return nil, error
		} else {
			user := model.User{Id: player.Id, LoginName: player.LoginName, Password: player.Password, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, YearsPlayed: player.YearsPlayed, Height: player.Height, Width: player.Width, Grip: player.Grip, Backhand: player.Backhand, Points: player.Points, IsAdult: player.IsAdult, CareerStats: player.CareerStats, Friends: friends, AllUnfinishedGames: unfinishedGames, AllHistoryGames: games, AllClubs: clubs, AllEvents: utils.IntMatrix{}, AllSchedules: schedules, AllOrders: allOrders, Addresss: addresss, Cart: cart, DefaultAddress: defaultAddress, Token: tokenStr}
			return &user, nil
		}
	} else {
		return nil, errors.New("no such user")
	}
}

func (impl *TennisMomentDaoImpl) ResetPassword(ctx context.Context, userLoginName string, userPassword string) bool {
	var Player model.Player
	impl.db.Where("login_name = ?", userLoginName).First(&Player)
	if Player.LoginName == userLoginName {
		player := model.Player{Password: userPassword}
		impl.db.Where("login_name = ?", Player.LoginName).Updates(&player)
		return true
	}
	return false
}

func (impl *TennisMomentDaoImpl) UpdatePlayer(ctx context.Context, player model.PlayerResponse) model.PlayerResponse {
	Player := model.Player{Id: player.Id, LoginName: player.LoginName, Password: player.Password, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, YearsPlayed: player.YearsPlayed, Height: player.Height, Width: player.Width, Grip: player.Grip, Backhand: player.Backhand, Points: player.Points, IsAdult: player.IsAdult, CareerStatsId: player.CareerStats.Id}
	impl.db.Where("id = ?", player.Id).Updates(&Player)
	impl.db.First(&Player, "id = ?", player.Id)
	stats := impl.UpdateCareerStats(ctx, player.CareerStats)
	return model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, YearsPlayed: Player.YearsPlayed, Height: Player.Height, Width: Player.Width, Grip: Player.Grip, Backhand: Player.Backhand, Points: Player.Points, IsAdult: Player.IsAdult, CareerStats: stats}
}

func (impl *TennisMomentDaoImpl) SearchPlayer(ctx context.Context, loginName string) bool {
	var Player model.Player
	impl.db.Where("login_name = ?", loginName).First(&Player)
	return Player.LoginName == loginName
}

func (impl *TennisMomentDaoImpl) GetPlayerInfo(ctx context.Context, id int64) (*model.PlayerResponse, error) {
	var Player model.Player
	impl.db.Where("id = ?", id).First(&Player)
	if Player.Id == id {
		stats := impl.GetStatsInfo(ctx, Player.CareerStatsId)
		return &model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, YearsPlayed: Player.YearsPlayed, Height: Player.Height, Width: Player.Width, Grip: Player.Grip, Backhand: Player.Backhand, Points: Player.Points, IsAdult: Player.IsAdult, CareerStats: stats}, nil
	}
	return nil, errors.New("no such player")
}

func (impl *TennisMomentDaoImpl) GetPlayerInfoByLoginName(ctx context.Context, loginName string) (*model.PlayerResponse, error) {
	var Player model.Player
	impl.db.Where("login_name = ?", loginName).First(&Player)
	if Player.LoginName == loginName {
		stats := impl.GetStatsInfo(ctx, Player.CareerStatsId)
		return &model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, YearsPlayed: Player.YearsPlayed, Height: Player.Height, Width: Player.Width, Grip: Player.Grip, Backhand: Player.Backhand, Points: Player.Points, IsAdult: Player.IsAdult, CareerStats: stats}, nil
	}
	return nil, errors.New("no such player")
}

func (impl *TennisMomentDaoImpl) GetPlayerInfoByLoginNameAndPassword(ctx context.Context, loginName string, password string) (*model.PlayerResponse, error) {
	var Player model.Player
	fmt.Println(password)

	impl.db.Where("login_name = ?", loginName).First(&Player)
	if Player.LoginName == loginName {
		// 将密码转换为字节数组
		passwordBytes := []byte(Player.Password)

		// 计算 SHA256 哈希值
		hash := sha256.Sum256(passwordBytes)

		// 将哈希值转换为字符串并打印输出
		authedPassword := hex.EncodeToString(hash[:])
		fmt.Println(authedPassword)
		if authedPassword == password {
			stats := impl.GetStatsInfo(ctx, Player.CareerStatsId)
			return &model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, YearsPlayed: Player.YearsPlayed, Height: Player.Height, Width: Player.Width, Grip: Player.Grip, Backhand: Player.Backhand, Points: Player.Points, IsAdult: Player.IsAdult, CareerStats: stats}, nil
		} else {
			return nil, errors.New("no such account or password")
		}
	}
	return nil, errors.New("no such user")
}

func (impl *TennisMomentDaoImpl) GetAllFriends(ctx context.Context, id int64) []model.PlayerResponse {
	var friends []model.Player
	var friendResponses []model.PlayerResponse
	var relationship1 []model.Relationship
	impl.db.Find(&relationship1, "player1_id", id)
	for _, friend := range relationship1 {
		if friend.Player1Id != friend.Player2Id {
			var player model.Player
			impl.db.First(&player, "id", friend.Player2Id)
			friends = append(friends, player)
		}
	}
	var relationship2 []model.Relationship
	impl.db.Find(&relationship2, "player2_id", id)
	for _, friend := range relationship2 {
		var player model.Player
		impl.db.First(&player, "id", friend.Player1Id)
		friends = append(friends, player)
	}

	for _, friend := range friends {
		stats := impl.GetStatsInfo(ctx, friend.CareerStatsId)
		friendResponse := model.PlayerResponse{Id: friend.Id, LoginName: friend.LoginName, Password: friend.Password, Name: friend.Name, Icon: friend.Icon, Sex: friend.Sex, Age: friend.Age, YearsPlayed: friend.YearsPlayed, Height: friend.Height, Width: friend.Width, Grip: friend.Grip, Backhand: friend.Backhand, Points: friend.Points, IsAdult: friend.IsAdult, CareerStats: stats}
		friendResponses = append(friendResponses, friendResponse)
	}
	return friendResponses
}

func (impl *TennisMomentDaoImpl) AddFriend(ctx context.Context, relationship model.Relationship) []model.PlayerResponse {
	res := impl.SearchFriend(ctx, relationship)
	if !res {
		impl.db.Create(&relationship)
	}
	return impl.GetAllFriends(ctx, relationship.Player1Id)
}

func (impl *TennisMomentDaoImpl) UpdateFriend(ctx context.Context, relationship model.Relationship) []model.PlayerResponse {
	res1 := impl.SearchFriend(ctx, model.Relationship{Player1Id: relationship.Player2Id, Player2Id: relationship.Player1Id})
	res2 := impl.SearchFriend(ctx, relationship)
	if !res1 && !res2 {
		impl.db.Create(&relationship)
	}

	return impl.GetAllFriends(ctx, relationship.Player1Id)
}

func (impl *TennisMomentDaoImpl) UpdateFriendDB(ctx context.Context, userId int64, friends []model.PlayerResponse) []model.PlayerResponse {
	relationships := make([]model.Relationship, 0)
	for _, friend := range friends {
		relationship := model.Relationship{Player1Id: userId, Player2Id: friend.Id}
		relationships = append(relationships, relationship)
	}
	var results []model.Relationship

	m := make(map[model.Relationship]bool)
	impl.db.Where("player1_id = ? OR player2_id = ?", userId, userId).Find(&results)
	for _, relat := range relationships {
		rerat := model.Relationship{Player1Id: relat.Player2Id, Player2Id: relat.Player1Id}
		m[relat] = true
		m[rerat] = true
	}
	for _, res := range results {
		rerat := model.Relationship{Player1Id: res.Player2Id, Player2Id: res.Player1Id}
		if !m[res] && !m[rerat] {
			impl.db.Where("player1_id = ? AND player2_id = ?", res.Player1Id, res.Player2Id).Delete(&res)
			impl.db.Where("player1_id = ? AND player2_id = ?", rerat.Player1Id, rerat.Player2Id).Delete(&rerat)
		}
	}
	return impl.GetAllFriends(ctx, userId)
}

func (impl *TennisMomentDaoImpl) AddSchedule(ctx context.Context, schedule model.Schedule) []model.ScheduleResponse {
	newSchedule := model.Schedule{}
	impl.db.Create(&newSchedule)
	newSchedule.Place = schedule.Place
	newSchedule.StartDate = schedule.StartDate
	newSchedule.Player1Id = schedule.Player1Id
	newSchedule.Player2Id = schedule.Player2Id
	impl.db.Save(&newSchedule)
	return impl.GetSchedules(ctx, schedule.Player1Id)
}

func (impl *TennisMomentDaoImpl) DeleteFriend(ctx context.Context, relationship model.Relationship) []model.PlayerResponse {
	if impl.SearchFriend(ctx, relationship) {
		impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player1Id, relationship.Player2Id).Delete(&relationship)
		impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player2Id, relationship.Player1Id).Delete(&relationship)
	}
	return impl.GetAllFriends(ctx, relationship.Player1Id)
}

func (impl *TennisMomentDaoImpl) SearchFriend(ctx context.Context, relationship model.Relationship) bool {
	var Relationship1 model.Relationship
	var Relationship2 model.Relationship
	impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player1Id, relationship.Player2Id).First(&Relationship1)
	impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player2Id, relationship.Player1Id).First(&Relationship2)
	return Relationship1 == relationship || Relationship2 == relationship

}

func (impl *TennisMomentDaoImpl) AddPlayerStats(ctx context.Context, PlayerStats model.PlayerStats) bool {
	impl.db.Save(&PlayerStats)
	return true
}

func (impl *TennisMomentDaoImpl) GetStatsInfo(ctx context.Context, statsId int64) model.Stats {
	var Stats model.Stats
	impl.db.First(&Stats, "id", statsId)
	return Stats
}

func (impl *TennisMomentDaoImpl) UpdateStats(ctx context.Context, Stats model.Stats, careerStatsId int64) model.Stats {
	impl.db.Save(&Stats)
	careerStats := model.Stats{}
	impl.db.First(&careerStats, "id", careerStatsId)
	careerStats.Aces += Stats.Aces
	careerStats.DoubleFaults += Stats.DoubleFaults
	careerStats.ServePoints += Stats.ServePoints
	careerStats.FirstServePoints += Stats.FirstServePoints
	careerStats.FirstServePointsIn += Stats.FirstServePointsIn
	careerStats.FirstServePointsWon += Stats.FirstServePointsWon
	careerStats.SecondServePointsIn += Stats.SecondServePointsIn
	careerStats.SecondServePointsWon += Stats.SecondServePointsWon
	careerStats.BreakPointsFaced += Stats.BreakPointsFaced
	careerStats.BreakPointsSaved += Stats.BreakPointsSaved
	careerStats.ServeGamesPlayed += Stats.ServeGamesPlayed
	careerStats.ServeGamesWon += Stats.ServeGamesWon
	careerStats.ReturnAces += Stats.ReturnAces
	careerStats.ReturnServePoints += Stats.ReturnServePoints
	careerStats.FirstServeReturnPoints += Stats.FirstServeReturnPoints
	careerStats.FirstServeReturnPointsIn += Stats.FirstServeReturnPointsIn
	careerStats.FirstServeReturnPointsWon += Stats.FirstServeReturnPointsWon
	careerStats.SecondServeReturnPointsIn += Stats.SecondServeReturnPointsIn
	careerStats.SecondServeReturnPointsWon += Stats.SecondServeReturnPointsWon
	careerStats.BreakPointsOpportunities += Stats.BreakPointsOpportunities
	careerStats.BreakPointsConverted += Stats.BreakPointsConverted
	careerStats.ReturnGamesPlayed += Stats.ReturnGamesPlayed
	careerStats.ReturnGamesWon += Stats.ReturnGamesWon
	careerStats.NetPoints += Stats.NetPoints
	careerStats.UnforcedErrors += Stats.UnforcedErrors
	careerStats.ForehandWinners += Stats.ForehandWinners
	careerStats.BackhandWinners += Stats.BackhandWinners
	impl.db.Save(&careerStats)
	return Stats
}

func (impl *TennisMomentDaoImpl) UpdateCareerStats(ctx context.Context, careerStats model.Stats) model.Stats {
	impl.db.Where("id = ?", careerStats.Id).Updates(&careerStats)
	return careerStats
}

func (impl *TennisMomentDaoImpl) GetMyClubs(ctx context.Context, playerId int64) utils.IntMatrix {
	clubMembers := make([]model.ClubMember, 0)
	clubResponses := utils.IntMatrix{}
	impl.db.Where("member_id = ?", playerId).Find(&clubMembers)

	for _, ClubMember := range clubMembers {
		clubResponses = append(clubResponses, ClubMember.ClubId)
	}
	return clubResponses
}

func (impl *TennisMomentDaoImpl) GetClubInfo(ctx context.Context, clubId int64) model.ClubResponse {
	club := model.Club{}
	impl.db.Where("id = ?", clubId).Find(&club)
	owner, _ := impl.GetPlayerInfo(ctx, club.OwnerId)
	events := impl.GetAllEventsForClub(ctx, club.Id)
	clubResponse := model.ClubResponse{Id: club.Id, Icon: club.Icon, Name: club.Name, Intro: club.Intro, Owner: *owner, Address: club.Address, Events: events}

	return clubResponse
}

func (impl *TennisMomentDaoImpl) UpdateClub(ctx context.Context, ClubMember model.ClubMember) {
	clubMember := model.ClubMember{}
	impl.db.Where("club_id = ? And member_id = ?", ClubMember.ClubId, ClubMember.MemberId).First(&clubMember)
	if clubMember.ClubId != ClubMember.ClubId || clubMember.MemberId != ClubMember.MemberId {
		impl.db.Create(ClubMember)
	}
	return
}

func (impl *TennisMomentDaoImpl) UpdateClubDB(ctx context.Context, userId int64, clubs utils.IntMatrix) utils.IntMatrix {

	var results utils.IntMatrix

	m := make(map[int64]bool)
	impl.db.Where("member_id = ?", userId).Find(&results)

	for _, item := range clubs {
		m[item] = true
	}

	for _, item := range results {
		if !m[item] {
			club := model.ClubMember{ClubId: item, MemberId: userId}
			impl.db.Where("club_id = ?", item).Delete(&club)
		}
	}
	return impl.GetMyClubs(ctx, userId)
}

func (impl *TennisMomentDaoImpl) GetAllEventsForClub(ctx context.Context, clubId int64) []model.EventResponse {
	clubEvents := make([]model.ClubEvent, 0)
	eventResponses := make([]model.EventResponse, 0)
	impl.db.Where("club_id = ?", clubId).Find(&clubEvents)
	for _, clubEvent := range clubEvents {
		eventResponse := impl.GetEventInfo(ctx, clubEvent.EventId)
		eventResponses = append(eventResponses, eventResponse)
	}
	return eventResponses
}

func (impl *TennisMomentDaoImpl) GetEventInfo(ctx context.Context, eventId int64) model.EventResponse {
	event := model.Event{}
	impl.db.Where("id = ?", eventId).First(&event)
	draw := impl.GetDraw(ctx, event.Draw)
	schedule := impl.GetEventSchedule(ctx, event.Schedule)
	eventResponse := model.EventResponse{Id: event.Id, Icon: event.Icon, Name: event.Name, StartDate: event.StartDate, EndDate: event.EndDate, Level: event.Level, Draw: draw, Schedule: schedule}
	return eventResponse
}

func (impl *TennisMomentDaoImpl) GetDraw(ctx context.Context, draw utils.IntMatrix) []model.PlayerResponse {
	playerResponses := make([]model.PlayerResponse, 0)
	for _, player := range draw {
		playerResponse, _ := impl.GetPlayerInfo(ctx, player)
		playerResponses = append(playerResponses, *playerResponse)
	}
	return playerResponses
}

func (impl *TennisMomentDaoImpl) GetEventSchedule(ctx context.Context, schedule utils.IntMatrix2) [][]model.GameResponse {
	gameResponses := make([][]model.GameResponse, 0)
	for _, daySchdule := range schedule {
		daySchduleResponse := make([]model.GameResponse, 0)
		for _, game := range daySchdule {
			gameResponse := impl.GetGameInfo(ctx, game)
			daySchduleResponse = append(daySchduleResponse, gameResponse)
		}
		gameResponses = append(gameResponses, daySchduleResponse)
	}
	return gameResponses
}

func (impl *TennisMomentDaoImpl) GetSchedules(ctx context.Context, playerId int64) []model.ScheduleResponse {
	schedules := make([]model.Schedule, 0)
	scheduleResponses := make([]model.ScheduleResponse, 0)
	impl.db.Where("player1_id = ? or player2_id = ?", playerId, playerId).Find(&schedules)
	for _, schdule := range schedules {
		if schdule.Player1Id == playerId {
			opponent, _ := impl.GetPlayerInfo(ctx, schdule.Player2Id)
			scheduleResponse := model.ScheduleResponse{Id: schdule.Id, Place: schdule.Place, StartDate: schdule.StartDate, Opponent: *opponent}
			scheduleResponses = append(scheduleResponses, scheduleResponse)
		} else {
			opponent, _ := impl.GetPlayerInfo(ctx, schdule.Player1Id)
			scheduleResponse := model.ScheduleResponse{Place: schdule.Place, StartDate: schdule.StartDate, Opponent: *opponent}
			scheduleResponses = append(scheduleResponses, scheduleResponse)
		}
	}
	return scheduleResponses
}

func (impl *TennisMomentDaoImpl) UpdateEvent(ctx context.Context, EventPlayer model.EventPlayer) {
	eventPlayer := model.EventPlayer{}
	impl.db.Where("event_id = ? And player_id = ?", EventPlayer.EventId, EventPlayer.PlayerId).First(&eventPlayer)
	if eventPlayer.EventId != EventPlayer.EventId || eventPlayer.PlayerId != EventPlayer.PlayerId {
		impl.db.Create(EventPlayer)
	}
}

func (impl *TennisMomentDaoImpl) UpdateSchedule(ctx context.Context, Schedule model.Schedule) {
	schedule := model.Schedule{}
	impl.db.Where("id = ?", Schedule.Id).First(&schedule)
	if schedule.Id == Schedule.Id {
		impl.db.Save(&Schedule)
	} else {
		impl.db.Create(&Schedule)
	}
}

func (impl *TennisMomentDaoImpl) UpdateScheduleDB(ctx context.Context, userId int64, schedules []model.ScheduleResponse) []model.ScheduleResponse {

	var results []model.Schedule

	m := make(map[int64]bool)
	impl.db.Where("player1_id = ? OR player2_id = ?", userId, userId).Find(&results)
	fmt.Println(len(schedules))
	fmt.Println(len(results))
	for _, relat := range schedules {
		m[relat.Id] = true
	}
	for _, res := range results {
		if !m[res.Id] {
			impl.db.Where("id = ?", res.Id).Delete(&res)
		}
	}
	return impl.GetSchedules(ctx, userId)
}

func (impl *TennisMomentDaoImpl) UpdateOrder(ctx context.Context, Order model.OrderResponse) model.OrderResponse {
	order := model.Order{Id: Order.Id, ShippingAddressId: Order.ShippingAddress.Id, PlayerId: Order.PlayerId, State: Order.State, Payment: Order.Payment, CreatedTime: Order.CreatedTime, PayedTime: Order.PayedTime, CompletedTime: Order.CompletedTime}
	impl.db.Save(&order)
	return impl.GetOrderInfo(ctx, order.Id)
}

func (impl *TennisMomentDaoImpl) DeleteOrder(ctx context.Context, id int64) bool {

	order := model.Order{}
	impl.db.Where("id = ?", id).Delete(&order)
	bill := model.Bill{}
	impl.db.Where("order_id", id).Delete(&bill)

	return true
}
func (impl *TennisMomentDaoImpl) GetOrderInfo(ctx context.Context, orderId int64) model.OrderResponse {
	order := model.Order{}
	impl.db.Where("id = ?", orderId).First(&order)

	bills := impl.GetBillInfos(ctx, order.Id)
	shippingAddress := impl.GetAddressInfo(ctx, order.ShippingAddressId)

	return model.OrderResponse{Id: order.Id, Bills: bills, ShippingAddress: shippingAddress, CreatedTime: order.CreatedTime, PayedTime: order.PayedTime, CompletedTime: order.CompletedTime, State: order.State}
}

func (impl *TennisMomentDaoImpl) GetOrderInfosByUserId(ctx context.Context, playerId int64) []model.OrderResponse {
	orders := make([]model.Order, 0)
	impl.db.Where("player_id = ?", playerId).Find(&orders)

	orderResponses := make([]model.OrderResponse, 0)
	for _, order := range orders {
		bills := impl.GetBillInfos(ctx, order.Id)
		shippingAddress := impl.GetAddressInfo(ctx, order.ShippingAddressId)

		orderResponse := model.OrderResponse{Id: order.Id, Bills: bills, ShippingAddress: shippingAddress, CreatedTime: order.CreatedTime, PayedTime: order.PayedTime, CompletedTime: order.CompletedTime, State: order.State}

		orderResponses = append(orderResponses, orderResponse)
	}
	fmt.Println(orderResponses)
	return orderResponses
}

func (impl *TennisMomentDaoImpl) AddBillToCart(ctx context.Context, bill *model.Bill) model.OrderResponse {

	Bill := model.Bill{}
	impl.db.Where("com_id = ? AND option_id = ? AND order_id = ?", bill.ComId, bill.OptionId, bill.OrderId).First(&Bill)
	newBill := model.Bill{}
	cart := impl.GetOrderInfo(ctx, bill.OrderId)
	if Bill.ComId == bill.ComId && Bill.OptionId == bill.OptionId && Bill.OrderId == bill.OrderId {
		Bill.Quantity += bill.Quantity
		impl.db.Save(&Bill)
	} else {
		impl.db.Create(&newBill)
		bill.Id = newBill.Id
		impl.db.Save(&bill)
		billInfo := impl.GetBillInfo(ctx, bill.Id)
		cart.Bills = append(cart.Bills, billInfo)
	}

	return *&cart
}

func (impl *TennisMomentDaoImpl) DeleteBillInCart(ctx context.Context, bill *model.Bill) model.OrderResponse {

	if impl.SearchBill(ctx, bill.Id) {
		impl.db.Where("id = ?", bill.Id).Delete(&bill)
	}

	cart := impl.GetOrderInfo(ctx, bill.OrderId)
	return cart
}

func (impl *TennisMomentDaoImpl) AssignCartForUser(ctx context.Context, order model.OrderResponse) int64 {

	Order := model.Order{Id: order.Id, ShippingAddressId: order.ShippingAddress.Id, PlayerId: order.PlayerId, State: order.State, Payment: order.Payment, CreatedTime: order.CreatedTime, PayedTime: order.PayedTime, CompletedTime: order.CompletedTime}
	impl.db.Save(&Order)
	for _, bill := range order.Bills {
		billResponse := model.Bill{Id: bill.Id, ComId: bill.Com.Id, Quantity: bill.Quantity, OptionId: bill.Option.Id, OrderId: order.Id}
		impl.AddBill(ctx, billResponse)
	}
	cart := impl.AddCartForPlayer(ctx, order.PlayerId)

	return cart

}

func (impl *TennisMomentDaoImpl) GetAllCommodities(ctx context.Context) []model.CommodityResponse {
	commodities := make([]model.Commodity, 0)
	impl.db.Find(&commodities)

	CommodityResponses := make([]model.CommodityResponse, 0)
	for _, commodity := range commodities {
		CommodityResponse := impl.GetCommodityInfo(ctx, commodity.Id)
		CommodityResponses = append(CommodityResponses, CommodityResponse)
	}

	return CommodityResponses
}

func (impl *TennisMomentDaoImpl) GetAllOrders(ctx context.Context) []model.OrderResponse {
	orders := make([]model.Order, 0)
	impl.db.Find(&orders)

	orderResponses := make([]model.OrderResponse, 0)
	for _, orderResponse := range orders {
		orderResponse := impl.GetOrderInfo(ctx, orderResponse.Id)
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses
}

func (impl *TennisMomentDaoImpl) GetBillInfos(ctx context.Context, orderId int64) []model.BillResponse {
	bills := make([]model.Bill, 0)
	impl.db.Where("order_id = ?", orderId).Find(&bills)

	billResponses := make([]model.BillResponse, 0)
	for _, bill := range bills {
		commodity := impl.GetCommodityInfo(ctx, bill.ComId)
		option := impl.GetOptionInfo(ctx, bill.OptionId)
		billResponse := model.BillResponse{Id: bill.Id, Com: commodity, Quantity: bill.Quantity, Option: option}
		billResponses = append(billResponses, billResponse)
	}

	return billResponses
}

func (impl *TennisMomentDaoImpl) GetOptionInfo(ctx context.Context, optionId int64) model.OptionResponse {
	option := model.Option{}
	impl.db.Where("id = ?", optionId).First(&option)

	return model.OptionResponse{Id: option.Id, Image: option.Image, Intro: option.Intro, Price: option.Price, Inventory: option.Inventory}
}

func (impl *TennisMomentDaoImpl) GetAddressInfo(ctx context.Context, addressId int64) model.AddressResponse {
	address := model.Address{}
	impl.db.Where("id = ?", addressId).First(&address)
	addressResponse := model.AddressResponse{Id: address.Id, Name: address.Name, Sex: address.Sex, PhoneNumber: address.PhoneNumber, Province: address.Province, City: address.City, Area: address.Area, DetailAddress: address.DetailedAddress, IsDefault: address.IsDefault}
	return addressResponse
}

func (impl *TennisMomentDaoImpl) AddOrder(ctx context.Context, Order model.OrderResponse) int64 {

	newOrder := model.Order{}
	impl.db.Create(&newOrder)
	order := model.Order{Id: newOrder.Id, ShippingAddressId: Order.ShippingAddress.Id, PlayerId: Order.PlayerId, State: Order.State, Payment: Order.Payment, CreatedTime: Order.CreatedTime, PayedTime: Order.PayedTime, CompletedTime: Order.CompletedTime}
	impl.db.Save(&order)
	for _, bill := range Order.Bills {
		billResponse := model.Bill{Id: bill.Id, ComId: bill.Com.Id, Quantity: bill.Quantity, OptionId: bill.Option.Id, OrderId: order.Id}
		impl.AddBill(ctx, billResponse)
	}
	return order.Id
}

func (impl *TennisMomentDaoImpl) SearchOrder(ctx context.Context, id int64) bool {
	var Order model.Order
	impl.db.Where("id = ?", id).First(&Order)
	return Order.Id == id
}

func (impl *TennisMomentDaoImpl) AddAddress(ctx context.Context, address model.Address) model.AddressResponse {

	newAddress := model.Address{}
	impl.db.Create(&newAddress)
	address.Id = newAddress.Id
	impl.db.Save(&address)
	return model.AddressResponse{Id: address.Id, Name: address.Name, Sex: address.Sex, PhoneNumber: address.PhoneNumber, Province: address.Province, City: address.City, Area: address.Area, DetailAddress: address.DetailedAddress, IsDefault: address.IsDefault}
}

func (impl *TennisMomentDaoImpl) SearchAddress(ctx context.Context, id int64) bool {
	var Address model.Address
	impl.db.Where("id = ?", id).First(&Address)
	return Address.Id == id
}

func (impl *TennisMomentDaoImpl) UpdateAddress(ctx context.Context, updatingAddress model.Address) model.AddressResponse {
	impl.db.Save(&updatingAddress)
	addresss := impl.GetMyAddresss(ctx, updatingAddress.PlayerId)
	for _, address := range addresss {
		if address != updatingAddress.Id {
			addressInfo := impl.GetAddressInfo(ctx, address)
			addressInfo.IsDefault = false
			AddressInfo := model.Address{Id: addressInfo.Id, PlayerId: updatingAddress.PlayerId, Name: addressInfo.Name, PhoneNumber: addressInfo.PhoneNumber, Sex: addressInfo.Sex, Province: addressInfo.Province, City: addressInfo.City, Area: addressInfo.Area, DetailedAddress: addressInfo.DetailAddress, IsDefault: addressInfo.IsDefault}
			impl.db.Save(&AddressInfo)
		}
	}
	return impl.GetAddressInfo(ctx, updatingAddress.Id)
}

func (impl *TennisMomentDaoImpl) DeleteAddress(ctx context.Context, id int64) bool {
	address := model.Address{}
	impl.db.Where("id = ?", id).Delete(&address)
	return true
}

func (impl *TennisMomentDaoImpl) AddBill(ctx context.Context, bill model.Bill) model.Bill {

	newBill := model.Bill{}
	impl.db.Create(&newBill)
	bill.Id = newBill.Id
	impl.db.Save(&bill)
	return bill
}

func (impl *TennisMomentDaoImpl) UpdateBill(ctx context.Context, userId int64, Bill model.Bill) model.BillResponse {
	impl.db.Save(&Bill)
	return impl.GetBillInfo(ctx, Bill.Id)
}

func (impl *TennisMomentDaoImpl) GetBillInfo(ctx context.Context, billId int64) model.BillResponse {
	bill := model.Bill{}
	impl.db.Where("id = ?", billId).First(&bill)
	com := impl.GetCommodityInfo(ctx, bill.ComId)
	option := impl.GetOptionInfo(ctx, bill.OptionId)
	billResponse := model.BillResponse{Id: bill.Id, Com: com, Quantity: bill.Quantity, Option: option}
	return billResponse
}
func (impl *TennisMomentDaoImpl) SearchBill(ctx context.Context, id int64) bool {
	var bill model.Bill
	impl.db.Where("id = ?", id).First(&bill)
	return bill.Id == id
}

func (impl *TennisMomentDaoImpl) AddCommodity(ctx context.Context, Commodity model.CommodityResponse) model.CommodityResponse {

	newCommodity := model.Commodity{}
	impl.db.Create(&newCommodity)
	newCommodity.Name = Commodity.Name
	newCommodity.Intro = Commodity.Intro
	newCommodity.Cag = Commodity.Cag
	newCommodity.State = Commodity.State
	for _, option := range Commodity.Options {
		impl.AddOption(ctx, option, newCommodity.Id)
	}
	impl.db.Save(&newCommodity)
	Commodity.Id = newCommodity.Id
	return Commodity
}

func (impl *TennisMomentDaoImpl) AddOption(ctx context.Context, option model.OptionResponse, comId int64) []model.OptionResponse {
	newOption := model.Option{}
	impl.db.Create(&newOption)
	newOption.Image = option.Image
	newOption.Intro = option.Intro
	newOption.Price = option.Price
	newOption.Inventory = option.Inventory
	newOption.ComId = comId
	println(newOption.ComId)
	impl.db.Save(&newOption)
	return impl.GetOptionsForCommidity(ctx, comId)
}

func (impl *TennisMomentDaoImpl) DeleteCommodity(ctx context.Context, id int64) []model.CommodityResponse {
	commodity := model.Commodity{}
	impl.db.Where("id = ?", id).Delete(&commodity)
	option := model.Option{}
	impl.db.Where("com_id = ?", id).Delete(&option)
	return impl.GetAllCommodities(ctx)
}

func (impl *TennisMomentDaoImpl) DeleteOption(ctx context.Context, id int64) []model.OptionResponse {
	Option := model.Option{}
	OptionRequest := model.Option{}
	impl.db.Where("id = ?", id).Find(&Option)
	impl.db.Where("id = ?", id).Delete(&OptionRequest)
	return impl.GetOptionsForCommidity(ctx, Option.ComId)
}

func (impl *TennisMomentDaoImpl) UpdateCommodity(ctx context.Context, CommodityResponse model.CommodityResponse) model.CommodityResponse {
	Commodity := model.Commodity{Id: CommodityResponse.Id, Name: CommodityResponse.Name, Intro: CommodityResponse.Intro, Cag: CommodityResponse.Cag, State: CommodityResponse.State}
	impl.db.Save(&Commodity)
	for _, option := range CommodityResponse.Options {
		impl.UpdateOption(ctx, option, CommodityResponse.Id)
	}
	return impl.GetCommodityInfo(ctx, Commodity.Id)
}

func (impl *TennisMomentDaoImpl) UpdateOption(ctx context.Context, OptionResponse model.OptionResponse, comId int64) []model.OptionResponse {
	Option := model.Option{Id: OptionResponse.Id, Image: OptionResponse.Image, Intro: OptionResponse.Intro, Price: OptionResponse.Price, Inventory: OptionResponse.Inventory, ComId: comId}
	impl.db.Save(&Option)
	return impl.GetOptionsForCommidity(ctx, comId)
}

func (impl *TennisMomentDaoImpl) GetCommodityInfo(ctx context.Context, CommodityId int64) model.CommodityResponse {
	Commodity := model.Commodity{}
	impl.db.Where("id = ?", CommodityId).First(&Commodity)
	orders := impl.GetOrderNumForCommodity(ctx, Commodity.Id)
	options := impl.GetOptionsForCommidity(ctx, Commodity.Id)
	return model.CommodityResponse{Id: Commodity.Id, Name: Commodity.Name, Intro: Commodity.Intro, Cag: Commodity.Cag, Orders: orders, Options: options, State: Commodity.State}
}

func (impl *TennisMomentDaoImpl) GetCommoditySimpleInfo(ctx context.Context, CommodityId int64) model.Commodity {
	Commodity := model.Commodity{}
	impl.db.Where("id = ?", CommodityId).First(&Commodity)
	return Commodity
}

func (impl *TennisMomentDaoImpl) GetOrderNumForCommodity(ctx context.Context, CommodityId int64) int64 {
	bills := make([]model.Bill, 0)
	impl.db.Where("id = ?", CommodityId).Find(&bills)
	var num int64
	for _, bill := range bills {
		num += bill.Quantity
	}
	return num
}

func (impl *TennisMomentDaoImpl) GetOptionsForCommidity(ctx context.Context, CommodityId int64) []model.OptionResponse {
	options := make([]model.Option, 0)
	impl.db.Where("com_id = ?", CommodityId).Find(&options)

	optionResponses := make([]model.OptionResponse, 0)
	for _, option := range options {
		optionResponse := model.OptionResponse{Id: option.Id, Image: option.Image, Intro: option.Intro, Price: option.Price, Inventory: option.Inventory}
		optionResponses = append(optionResponses, optionResponse)
	}
	return optionResponses
}

func (impl *TennisMomentDaoImpl) SearchCommodity(ctx context.Context, id int64) bool {
	var Commodity model.Commodity
	impl.db.Where("id = ?", id).First(&Commodity)
	return Commodity.Id == id
}

func (impl *TennisMomentDaoImpl) GetMyOrders(ctx context.Context, playerId int64) utils.IntMatrix {
	orders := make([]model.Order, 0)
	orderResponses := utils.IntMatrix{}
	impl.db.Where("player_id = ?", playerId).Find(&orders)

	for _, order := range orders {
		orderResponses = append(orderResponses, order.Id)
	}
	return orderResponses
}

func (impl *TennisMomentDaoImpl) GetMyAddresss(ctx context.Context, playerId int64) utils.IntMatrix {
	addresss := make([]model.Address, 0)
	addressResponses := utils.IntMatrix{}
	impl.db.Where("player_id = ?", playerId).Find(&addresss)

	for _, order := range addresss {
		addressResponses = append(addressResponses, order.Id)
	}
	return addressResponses
}
func (impl *TennisMomentDaoImpl) GetMyDefaultAddress(ctx context.Context, playerId int64) model.AddressResponse {
	address := model.Address{}
	impl.db.Where("player_id = ? AND is_default = ?", playerId, 1).Find(&address)
	addressResponse := model.AddressResponse{Id: address.Id, Name: address.Name, Sex: address.Sex, PhoneNumber: address.PhoneNumber, Province: address.Province, City: address.City, Area: address.Area, DetailAddress: address.DetailedAddress, IsDefault: address.IsDefault}
	return addressResponse
}

func (impl *TennisMomentDaoImpl) GetMyCart(ctx context.Context, playerId int64) int64 {
	cart := model.Cart{}
	impl.db.Where("player_id = ?", playerId).First(&cart)
	return cart.OrderId
}
