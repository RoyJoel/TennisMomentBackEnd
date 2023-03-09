package model

import (
	"encoding/json"

	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"
)

type Game struct {
	Date                float64         `json:"date"`
	Place               string          `json:"place"`
	Surface             string          `json:"surface"`
	SetNum              int             `json:"setNum"`
	GameNum             int             `json:"gameNum"`
	IsGoldenGoal        bool            `json:"isGoldenGoal"`
	IsPlayer1Serving    bool            `json:"isPlayer1Serving"`
	IsCompleted         bool            `json:"isCompleted"`
	Player1LoginName    string          `json:"player1LoginName"`
	Player1StatsId      int             `json: "player1Stats"`
	Player2LoginName    string          `json:"player2LoginName"`
	Player2StatsId      int             `json: "player2Stats"`
	IsPlayer1FirstServe bool            `json:"isPlayer1FirstServe"`
	IsPlayer2FirstServe bool            `json:"isPlayer2FirstServe"`
	Result              utils.IntMatrix `json:"result"`
}

func (game Game) TableName() string {
	return "Game"
}

func (game Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"date":                game.Date,
		"place":               game.Place,
		"surface":             game.Surface,
		"setNum":              game.SetNum,
		"gameNum":             game.GameNum,
		"isGoldenGoal":        game.IsGoldenGoal,
		"isPlayer1Serving":    game.IsPlayer1Serving,
		"isCompleted":         game.IsCompleted,
		"player1LoginName":    game.Player1LoginName,
		"player1StatsId":      game.Player1StatsId,
		"player2LoginName":    game.Player2LoginName,
		"player2StatsId":      game.Player2StatsId,
		"isPlayer1FirstServe": game.IsPlayer1FirstServe,
		"isPlayer2FirstServe": game.IsPlayer2FirstServe,
		"result":              game.Result,
	})
}

// Redis类似序列化操作
func (game Game) MarshalBinary() ([]byte, error) {
	return json.Marshal(game)
}

func (game Game) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &game)
}
