package model

import (
	"encoding/json"
)

type Player struct {
	Id            int64   `json:"id"`
	LoginName     string  `json:"loginName"`
	Name          string  `json:"name"`
	Icon          string  `json:"icon"`
	Sex           string  `json:"sex"`
	Age           int64   `json:"age"`
	YearsPlayed   int64   `json:"yearsPlayed"`
	Height        float32 `json:"height"`
	Width         float32 `json:"width"`
	Grip          string  `json:"grip"`
	Backhand      string  `json:"backhand"`
	Points        int64   `json:"points"`
	IsAdult       bool    `json:"isAdult"`
	CareerStatsId int64   `json:"careerStatsId,omitempty"`
}

func (player Player) TableName() string {
	return "Player"
}

func (player Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":            player.Id,
		"loginName":     player.LoginName,
		"name":          player.Name,
		"icon":          player.Icon,
		"sex":           player.Sex,
		"age":           player.Age,
		"yearsPlayed":   player.YearsPlayed,
		"height":        player.Height,
		"width":         player.Width,
		"grip":          player.Grip,
		"backhand":      player.Backhand,
		"points":        player.Points,
		"isAdult":       player.IsAdult,
		"careerStatsId": player.CareerStatsId,
	})
}

// Redis类似序列化操作
func (player Player) MarshalBinary() ([]byte, error) {
	return json.Marshal(player)
}

func (player Player) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &player)
}

type PlayerResponse struct {
	Id          int64   `json:"id"`
	LoginName   string  `json:"loginName"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
	Sex         string  `json:"sex"`
	Age         int64   `json:"age"`
	YearsPlayed int64   `json:"yearsPlayed"`
	Height      float32 `json:"height"`
	Width       float32 `json:"width"`
	Grip        string  `json:"grip"`
	Backhand    string  `json:"backhand"`
	Points      int64   `json:"points"`
	IsAdult     bool    `json:"isAdult"`
	CareerStats Stats   `json:"careerStats,omitempty"`
}

func (playerResponse PlayerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":          playerResponse.Id,
		"loginName":   playerResponse.LoginName,
		"name":        playerResponse.Name,
		"icon":        playerResponse.Icon,
		"sex":         playerResponse.Sex,
		"age":         playerResponse.Age,
		"yearsPlayed": playerResponse.YearsPlayed,
		"height":      playerResponse.Height,
		"width":       playerResponse.Width,
		"grip":        playerResponse.Grip,
		"backhand":    playerResponse.Backhand,
		"points":      playerResponse.Points,
		"isAdult":     playerResponse.IsAdult,
		"careerStats": playerResponse.CareerStats,
	})
}

// Redis类似序列化操作
func (playerResponse PlayerResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(playerResponse)
}

func (playerResponse PlayerResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &playerResponse)
}
