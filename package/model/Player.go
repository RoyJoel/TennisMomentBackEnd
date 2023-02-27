package model

import (
	"encoding/json"
)

type Player struct {
	Id          int64   `json:"id"`
	LoginName   string  `json:"loginName"`
	Name        string  `json:"name"`
	Icon        []byte  `json:"icon"`
	Sex         string  `json:"sex"`
	Age         int64   `json:"age"`
	YearsPlayed int64   `json:"yearsPlayed"`
	Height      float32 `json:"height"`
	Width       float32 `json:"width"`
	Grip        string  `json:"grip"`
	Backhand    string  `json:"backhand"`
	// friends     []Player `json:"friends"`
	// gamesPlayed []Game   `json:"gamesPlayed"`
}

func (player Player) TableName() string {
	return "Player"
}

func (player Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":          player.Id,
		"loginName":   player.LoginName,
		"name":        player.Name,
		"icon":        player.Icon,
		"sex":         player.Sex,
		"age":         player.Age,
		"yearsPlayed": player.YearsPlayed,
		"height":      player.Height,
		"width":       player.Width,
		"grip":        player.Grip,
		"backhand":    player.Backhand,
		// "friends":     player.friends,
		// "gamesPlayed": player.gamesPlayed,
	})
}

// Redis类似序列化操作
func (player Player) MarshalBinary() ([]byte, error) {
	return json.Marshal(player)
}

func (player Player) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &player)
}
