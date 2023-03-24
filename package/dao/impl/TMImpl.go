package impl

import (
	"context"
	"fmt"

	"github.com/RoyJoel/TennisMomentBackEnd/package/cache"
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
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

	player1 := impl.GetPlayerInfo(ctx, game.Player1Id)
	player2 := impl.GetPlayerInfo(ctx, game.Player2Id)

	impl.db.Create(&Stats1)
	impl.db.Create(&Stats2)

	PlayerStats1 := model.PlayerStats{PlayerId: game.Player1Id, StatsId: Stats1.Id}
	PlayerStats2 := model.PlayerStats{PlayerId: game.Player2Id, StatsId: Stats2.Id}

	impl.db.Create(&PlayerStats1)
	impl.db.Create(&PlayerStats2)

	game.Player1StatsId = Stats1.Id
	game.Player2StatsId = Stats2.Id

	impl.db.Create(&game)

	gameResponse = model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: player1, Player1Stats: Stats1, Player2: player2, Player2Stats: Stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
	return *&gameResponse
}

func (impl *TennisMomentDaoImpl) UpdateGameAndStats(ctx context.Context, gameResponse model.GameResponse) model.GameResponse {
	game := model.Game{Id: gameResponse.Id, Place: gameResponse.Place, Surface: gameResponse.Surface, SetNum: gameResponse.SetNum, GameNum: gameResponse.GameNum, Round: gameResponse.Round, IsGoldenGoal: gameResponse.IsGoldenGoal, IsPlayer1Serving: gameResponse.IsPlayer1Serving, IsPlayer1Left: gameResponse.IsPlayer1Left, IsChangePosition: gameResponse.IsChangePosition, StartDate: gameResponse.StartDate, EndDate: gameResponse.EndDate, Player1Id: gameResponse.Player1.Id, Player1StatsId: gameResponse.Player1Stats.Id, Player2Id: gameResponse.Player2.Id, Player2StatsId: gameResponse.Player2Stats.Id, IsPlayer1FirstServe: gameResponse.IsPlayer1FirstServe, IsPlayer2FirstServe: gameResponse.IsPlayer2FirstServe, Result: gameResponse.Result}

	impl.db.Where("id = ?", game.Id).Save(game)
	impl.UpdateStats(ctx, gameResponse.Player1Stats, gameResponse.Player1.CareerStats.Id)
	impl.UpdateStats(ctx, gameResponse.Player2Stats, gameResponse.Player2.CareerStats.Id)
	fmt.Println(gameResponse)
	return gameResponse
}

func (impl *TennisMomentDaoImpl) GetRecentGames(ctx context.Context, id int64, limit int) []model.GameResponse {
	var games []model.Game
	var gameResponses []model.GameResponse
	impl.db.Order("date desc").Limit(limit).Find(&games, "player1_id = ? OR player2_id = ?", id, id)
	for _, game := range games {
		player1 := impl.GetPlayerInfo(ctx, game.Player1Id)
		stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
		player2 := impl.GetPlayerInfo(ctx, game.Player2Id)
		stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
		gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: player1, Player1Stats: stats1, Player2: player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
		gameResponses = append(gameResponses, gameResponse)
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
		player1 := impl.GetPlayerInfo(ctx, game.Player1Id)
		stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
		player2 := impl.GetPlayerInfo(ctx, game.Player2Id)
		stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
		gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: player1, Player1Stats: stats1, Player2: player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
		gameResponses = append(gameResponses, gameResponse)
	}

	return gameResponses
}

func (impl *TennisMomentDaoImpl) GetGameInfo(ctx context.Context, id int64) model.GameResponse {
	var game model.Game
	impl.db.Where("id = ?", id).First(&game)
	if game.Id == id {
		player1 := impl.GetPlayerInfo(ctx, game.Player1Id)
		stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
		player2 := impl.GetPlayerInfo(ctx, game.Player2Id)
		stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
		return model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: player1, Player1Stats: stats1, Player2: player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
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
		player1 := impl.GetPlayerInfo(ctx, game.Player1Id)
		stats1 := impl.GetStatsInfo(ctx, game.Player1StatsId)
		player2 := impl.GetPlayerInfo(ctx, game.Player2Id)
		stats2 := impl.GetStatsInfo(ctx, game.Player2StatsId)
		gameResponse := model.GameResponse{Id: game.Id, Place: game.Place, Surface: game.Surface, SetNum: game.SetNum, GameNum: game.GameNum, Round: game.Round, IsGoldenGoal: game.IsGoldenGoal, IsPlayer1Serving: game.IsPlayer1Serving, IsPlayer1Left: game.IsPlayer1Left, IsChangePosition: game.IsChangePosition, StartDate: game.StartDate, EndDate: game.EndDate, Player1: player1, Player1Stats: stats1, Player2: player2, Player2Stats: stats2, IsPlayer1FirstServe: game.IsPlayer1FirstServe, IsPlayer2FirstServe: game.IsPlayer2FirstServe, Result: game.Result}
		gameResponses = append(gameResponses, gameResponse)
	}

	return gameResponses
}

func (impl *TennisMomentDaoImpl) AddPlayer(ctx context.Context, Player model.Player) model.PlayerResponse {

	res := impl.SearchPlayer(ctx, Player.Id)
	if !res {
		Stats := model.Stats{}
		impl.db.Create(&Stats)
		PlayerStats := model.PlayerStats{PlayerId: Player.Id, StatsId: Stats.Id}
		impl.db.Create(&PlayerStats)
		Player.CareerStatsId = Stats.Id
		impl.db.Create(&Player)
		return model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, YearsPlayed: Player.YearsPlayed, Height: Player.Height, Width: Player.Width, Grip: Player.Grip, Backhand: Player.Backhand, Points: Player.Points, IsAdult: Player.IsAdult, CareerStats: Stats}
	} else {
		var player model.Player
		impl.db.Where("id = ?", Player.Id).First(&player)
		stats := impl.GetStatsInfo(ctx, player.CareerStatsId)
		return model.PlayerResponse{Id: player.Id, LoginName: player.LoginName, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, YearsPlayed: player.YearsPlayed, Height: player.Height, Width: player.Width, Grip: player.Grip, Backhand: player.Backhand, Points: player.Points, IsAdult: player.IsAdult, CareerStats: stats}
	}
}

func (impl *TennisMomentDaoImpl) SignUp(ctx context.Context, user model.User) model.User {
	playerInfo := model.Player{Id: user.Id, LoginName: user.LoginName, Name: user.Name, Icon: user.Icon, Sex: user.Sex, Age: user.Age, YearsPlayed: user.YearsPlayed, Height: user.Height, Width: user.Width, Grip: user.Grip, Backhand: user.Backhand, Points: user.Points, IsAdult: user.IsAdult, CareerStatsId: user.CareerStats.Id}
	player := impl.AddPlayer(ctx, playerInfo)
	relationship := model.Relationship{Player1Id: player.Id, Player2Id: player.Id}
	impl.db.Create(relationship)
	friends := []model.PlayerResponse{}
	friends = append(friends, player)
	user = model.User{Id: player.Id, LoginName: player.LoginName, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, YearsPlayed: player.YearsPlayed, Height: player.Height, Width: player.Width, Grip: player.Grip, Backhand: player.Backhand, Points: player.Points, IsAdult: player.IsAdult, CareerStats: player.CareerStats, Friends: friends, AllGames: make([]model.GameResponse, 0), AllTourLevelGames: make([]model.GameResponse, 0), AllClubs: make([]model.ClubResponse, 0), AllEvents: make([]model.EventResponse, 0)}
	return user
}
func (impl *TennisMomentDaoImpl) SignIn(ctx context.Context, userId int64) model.User {
	player := impl.GetPlayerInfo(ctx, userId)
	friends := impl.GetAllFriends(ctx, player.Id)
	games := impl.GetAllGames(ctx, player.Id)
	clubs := impl.GetMyClubs(ctx, player.Id)
	fmt.Println(player.CareerStats)
	user := model.User{Id: player.Id, LoginName: player.LoginName, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, YearsPlayed: player.YearsPlayed, Height: player.Height, Width: player.Width, Grip: player.Grip, Backhand: player.Backhand, Points: player.Points, IsAdult: player.IsAdult, CareerStats: player.CareerStats, Friends: friends, AllGames: games, AllTourLevelGames: make([]model.GameResponse, 0), AllClubs: clubs, AllEvents: make([]model.EventResponse, 0)}
	return user
}

func (impl *TennisMomentDaoImpl) UpdatePlayer(ctx context.Context, player model.PlayerResponse) model.PlayerResponse {
	Player := model.Player{Id: player.Id, LoginName: player.LoginName, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, YearsPlayed: player.YearsPlayed, Height: player.Height, Width: player.Width, Grip: player.Grip, Backhand: player.Backhand, Points: player.Points, IsAdult: player.IsAdult, CareerStatsId: player.CareerStats.Id}
	impl.db.Where("id = ?", player.Id).Updates(Player)
	impl.db.First(&Player, "id = ?", player.Id)
	stats := impl.UpdateCareerStats(ctx, player.CareerStats)
	return model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, YearsPlayed: Player.YearsPlayed, Height: Player.Height, Width: Player.Width, Grip: Player.Grip, Backhand: Player.Backhand, Points: Player.Points, IsAdult: Player.IsAdult, CareerStats: stats}
}

func (impl *TennisMomentDaoImpl) SearchPlayer(ctx context.Context, id int64) bool {
	var Player model.Player
	impl.db.Where("id = ?", id).First(&Player)
	return Player.Id == id
}

func (impl *TennisMomentDaoImpl) GetPlayerInfo(ctx context.Context, id int64) model.PlayerResponse {
	var Player model.Player
	impl.db.Where("id = ?", id).First(&Player)
	if Player.Id == id {
		stats := impl.GetStatsInfo(ctx, Player.CareerStatsId)
		return model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, YearsPlayed: Player.YearsPlayed, Height: Player.Height, Width: Player.Width, Grip: Player.Grip, Backhand: Player.Backhand, Points: Player.Points, IsAdult: Player.IsAdult, CareerStats: stats}
	}
	return model.PlayerResponse{}
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
		friendResponse := model.PlayerResponse{Id: friend.Id, LoginName: friend.LoginName, Name: friend.Name, Icon: friend.Icon, Sex: friend.Sex, Age: friend.Age, YearsPlayed: friend.YearsPlayed, Height: friend.Height, Width: friend.Width, Grip: friend.Grip, Backhand: friend.Backhand, Points: friend.Points, IsAdult: friend.IsAdult, CareerStats: stats}
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

func (impl *TennisMomentDaoImpl) DeleteFriend(ctx context.Context, relationship model.Relationship) []model.PlayerResponse {
	if impl.SearchFriend(ctx, relationship) {
		impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player1Id, relationship.Player2Id).Delete(&relationship)
		impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player1Id, relationship.Player2Id).Delete(&relationship)
	}
	return impl.GetAllFriends(ctx, relationship.Player1Id)
}

func (impl *TennisMomentDaoImpl) SearchFriend(ctx context.Context, relationship model.Relationship) bool {
	var Relationship1 model.Relationship
	var Relationship2 model.Relationship
	impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player1Id, relationship.Player2Id).First(&Relationship1)
	impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player1Id, relationship.Player2Id).First(&Relationship2)
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
	impl.db.Save(Stats)
	careerStats := model.Stats{}
	impl.db.First(&careerStats, "id", careerStatsId)
	careerStats.Aces += Stats.Aces
	careerStats.DoubleFaults += Stats.DoubleFaults
	careerStats.ServePoints += Stats.ServePoints
	careerStats.FirstServePoints += Stats.FirstServePoints
	careerStats.FirstServePointsIn += Stats.FirstServePointsIn
	careerStats.FirstServePointsWon += Stats.FirstServePointsWon
	careerStats.SecondServePointsWon += Stats.SecondServePointsWon
	careerStats.BreakPointsFaced += Stats.BreakPointsFaced
	careerStats.BreakPointsSaved += Stats.BreakPointsSaved
	careerStats.ServeGamesPlayed += Stats.ServeGamesPlayed
	careerStats.ServeGamesWon += Stats.ServeGamesWon
	careerStats.ReturnAces += Stats.ReturnAces
	careerStats.ReturnServePoints += Stats.ReturnServePoints
	careerStats.FirstServeReturnPoints += Stats.FirstServeReturnPoints
	careerStats.FirstServeReturnPointsWon += Stats.FirstServeReturnPointsWon
	careerStats.SecondServeReturnPointsWon += Stats.SecondServeReturnPointsWon
	careerStats.BreakPointsOpportunities += Stats.BreakPointsOpportunities
	careerStats.BreakPointsConverted += Stats.BreakPointsConverted
	careerStats.ReturnGamesPlayed += Stats.ReturnGamesPlayed
	careerStats.ReturnGamesWon += Stats.ReturnGamesWon
	careerStats.NetPoints += Stats.NetPoints
	careerStats.UnforcedErrors += Stats.UnforcedErrors
	careerStats.ForehandWinners += Stats.ForehandWinners
	careerStats.BackhandWinners += Stats.BackhandWinners
	impl.db.Save(careerStats)
	return Stats
}

func (impl *TennisMomentDaoImpl) UpdateCareerStats(ctx context.Context, careerStats model.Stats) model.Stats {
	impl.db.Where("id = ?", careerStats.Id).Updates(careerStats)
	return careerStats
}

func (impl *TennisMomentDaoImpl) GetMyClubs(ctx context.Context, playerId int64) []model.ClubResponse {
	clubMembers := make([]model.ClubMember, 0)
	clubResponses := make([]model.ClubResponse, 0)
	impl.db.Where("member_id = ?", playerId).Find(&clubMembers)

	for _, ClubMember := range clubMembers {
		clubResponse := impl.GetClubInfo(ctx, ClubMember.ClubId)
		clubResponses = append(clubResponses, clubResponse)
	}
	return clubResponses
}

func (impl *TennisMomentDaoImpl) GetClubInfo(ctx context.Context, clubId int64) model.ClubResponse {
	club := model.Club{}
	impl.db.Where("id = ?", clubId).Find(&club)
	owner := impl.GetPlayerInfo(ctx, club.OwnerId)
	events := impl.GetAllEventsForClub(ctx, club.Id)
	clubResponse := model.ClubResponse{Id: club.Id, Icon: club.Icon, Name: club.Name, Intro: club.Intro, Owner: owner, Address: club.Address, Events: events}

	return clubResponse
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
	schedule := impl.GetSchedule(ctx, event.Schedule)
	eventResponse := model.EventResponse{Id: event.Id, Icon: event.Icon, Name: event.Name, StartDate: event.StartDate, EndDate: event.EndDate, Level: event.Level, Draw: draw, Schedule: schedule}
	return eventResponse
}

func (impl *TennisMomentDaoImpl) GetDraw(ctx context.Context, draw utils.IntMatrix) []model.PlayerResponse {
	playerResponses := make([]model.PlayerResponse, 0)
	for _, player := range draw {
		playerResponse := impl.GetPlayerInfo(ctx, player)
		playerResponses = append(playerResponses, playerResponse)
	}
	return playerResponses
}

func (impl *TennisMomentDaoImpl) GetSchedule(ctx context.Context, schedule utils.IntMatrix2) [][]model.GameResponse {
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
