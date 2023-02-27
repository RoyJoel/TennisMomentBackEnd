// package model

// import (
// 	"encoding/json"
// 	"time"
// )

// type Game struct {
// 	id          int64    `json:"id"`
// 	date        time.Time `json:"date"`
// 	place       string    `json:"place"`
// 	winner      Player    `json:"winner"`
// 	winnerStats Stats     `json:"winnerStats"`
// 	loser       Player    `json:"loser"`
// 	loserStats  Stats     `json:"loserStats"`
// }

// func (game Game) TableName() string {
// 	return "Game"
// }

// func (game Game) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(map[string]interface{}{
// 		"id":          game.id,
// 		"date":        game.date,
// 		"place":       game.place,
// 		"winner":      game.winner,
// 		"winnerStats": game.winnerStats,
// 		"loser":       game.loser,
// 		"loserStats":  game.loserStats,
// 	})
// }

// // Redis类似序列化操作
// func (game Game) MarshalBinary() ([]byte, error) {
// 	return json.Marshal(game)
// }

// func (game Game) UnmarshalBinary(data []byte) error {
// 	return json.Unmarshal(data, &game)
// }
package model