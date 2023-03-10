package model

import (
	"encoding/json"
)

type Player struct {
	LoginName     string       `json:"loginName"`
	Name          string       `json:"name"`
	Icon          string       `json:"icon"`
	Sex           string       `json:"sex"`
	Age           int64        `json:"age"`
	YearsPlayed   int64        `json:"yearsPlayed"`
	Height        float32      `json:"height"`
	Width         float32      `json:"width"`
	Grip          string       `json:"grip"`
	Backhand      string       `json:"backhand"`
	Points        int64        `json:"points"`
	IsAdult       bool         `json:"isAdult"`
	CareerStatsId int          `json:"careerStatsId,omitempty"`
	Friends       string       `json:"Friends,omitempty"`
	PlayerStats   PlayerStats  `gorm:"foreignKey:Friends;references:LoginName"`
	Relationship  Relationship `gorm:"foreignKey:Friends;references:LoginName"`
}

func (player Player) TableName() string {
	return "Player"
}

func (player Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
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
		"friends":       player.Friends,
	})
}

// Redis类似序列化操作
func (player Player) MarshalBinary() ([]byte, error) {
	return json.Marshal(player)
}

func (player Player) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &player)
}
