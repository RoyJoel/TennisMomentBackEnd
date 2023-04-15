package model

import (
	"encoding/json"

	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"
)

type Game struct {
	Id                  int64            `json:"id"`
	Place               string           `json:"place"`
	Surface             string           `json:"surface"`
	SetNum              int64            `json:"setNum"`
	GameNum             int64            `json:"gameNum"`
	Round               int64            `json:"round"`
	IsGoldenGoal        bool             `json:"isGoldenGoal"`
	IsPlayer1Serving    bool             `json:"isPlayer1Serving"`
	IsPlayer1Left       bool             `json:"isPlayer1Left"`
	IsChangePosition    bool             `json:"isChangePosition"`
	StartDate           float64          `json:"startDate"`
	EndDate             float64          `json:"endDate"`
	Player1Id           int64            `json:"player1Id"`
	Player1StatsId      int64            `json: "player1StatsId"`
	Player2Id           int64            `json:"player2Id"`
	Player2StatsId      int64            `json: "player2StatsId"`
	IsPlayer1FirstServe bool             `json:"isPlayer1FirstServe"`
	IsPlayer2FirstServe bool             `json:"isPlayer2FirstServe"`
	Result              utils.IntMatrix3 `json:"result"`
}

func (game Game) TableName() string {
	return "Game"
}

func (game Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                  game.Id,
		"place":               game.Place,
		"surface":             game.Surface,
		"setNum":              game.SetNum,
		"gameNum":             game.GameNum,
		"round":               game.Round,
		"isGoldenGoal":        game.IsGoldenGoal,
		"isPlayer1Serving":    game.IsPlayer1Serving,
		"isPlayer1Left":       game.IsPlayer1Left,
		"isChangePosition":    game.IsChangePosition,
		"startDate":           game.StartDate,
		"endDate":             game.EndDate,
		"player1Id":           game.Player1Id,
		"player1StatsId":      game.Player1StatsId,
		"player2Id":           game.Player2Id,
		"player2StatsId":      game.Player2StatsId,
		"isPlayer1FirstServe": game.IsPlayer1FirstServe,
		"isPlayer2FirstServe": game.IsPlayer2FirstServe,
		"result":              game.Result,
	})
}

func (g Game) Equals(other Game) bool {
	return g.Id == other.Id &&
		g.Place == other.Place &&
		g.Surface == other.Surface &&
		g.SetNum == other.SetNum &&
		g.GameNum == other.GameNum &&
		g.Round == other.Round &&
		g.IsGoldenGoal == other.IsGoldenGoal &&
		g.IsPlayer1Serving == other.IsPlayer1Serving &&
		g.IsPlayer1Left == other.IsPlayer1Left &&
		g.IsChangePosition == other.IsChangePosition &&
		g.StartDate == other.StartDate &&
		g.EndDate == other.EndDate &&
		g.Player1Id == other.Player1Id &&
		g.Player1StatsId == other.Player1StatsId &&
		g.Player2Id == other.Player2Id &&
		g.Player2StatsId == other.Player2StatsId &&
		g.IsPlayer1FirstServe == other.IsPlayer1FirstServe &&
		g.IsPlayer2FirstServe == other.IsPlayer2FirstServe &&
		g.Result.Equals(other.Result)
}

// Redis类似序列化操作
func (game Game) MarshalBinary() ([]byte, error) {
	return json.Marshal(game)
}

func (game Game) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &game)
}

type GameResponse struct {
	Id                  int64            `json:"id"`
	Place               string           `json:"place"`
	Surface             string           `json:"surface"`
	SetNum              int64            `json:"setNum"`
	GameNum             int64            `json:"gameNum"`
	Round               int64            `json:"round"`
	IsGoldenGoal        bool             `json:"isGoldenGoal"`
	IsPlayer1Serving    bool             `json:"isPlayer1Serving"`
	IsPlayer1Left       bool             `json:"isPlayer1Left"`
	IsChangePosition    bool             `json:"isChangePosition"`
	StartDate           float64          `json:"startDate"`
	EndDate             float64          `json:"endDate"`
	Player1             PlayerResponse   `json:"player1"`
	Player1Stats        Stats            `json: "player1Stats"`
	Player2             PlayerResponse   `json:"player2"`
	Player2Stats        Stats            `json: "player2Stats"`
	IsPlayer1FirstServe bool             `json:"isPlayer1FirstServe"`
	IsPlayer2FirstServe bool             `json:"isPlayer2FirstServe"`
	Result              utils.IntMatrix3 `json:"result"`
}

func (gameResponse GameResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                  gameResponse.Id,
		"place":               gameResponse.Place,
		"surface":             gameResponse.Surface,
		"setNum":              gameResponse.SetNum,
		"gameNum":             gameResponse.GameNum,
		"round":               gameResponse.Round,
		"isGoldenGoal":        gameResponse.IsGoldenGoal,
		"isPlayer1Serving":    gameResponse.IsPlayer1Serving,
		"isPlayer1Left":       gameResponse.IsPlayer1Left,
		"isChangePosition":    gameResponse.IsChangePosition,
		"startDate":           gameResponse.StartDate,
		"endDate":             gameResponse.EndDate,
		"player1":             gameResponse.Player1,
		"player1Stats":        gameResponse.Player1Stats,
		"player2":             gameResponse.Player2,
		"player2Stats":        gameResponse.Player2Stats,
		"isPlayer1FirstServe": gameResponse.IsPlayer1FirstServe,
		"isPlayer2FirstServe": gameResponse.IsPlayer2FirstServe,
		"result":              gameResponse.Result,
	})
}

// Redis类似序列化操作
func (gameResponse GameResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(gameResponse)
}

func (gameResponse GameResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &gameResponse)
}
