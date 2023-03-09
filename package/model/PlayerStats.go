package model

import (
	"encoding/json"
)

type PlayerStats struct {
	LoginName string `json:"loginName"`
	Id        int    `json:"id"`
}

func (PlayerStats PlayerStats) TableName() string {
	return "PlayerStats"
}

func (PlayerStats PlayerStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"loginName": PlayerStats.LoginName,
		"id":        PlayerStats.Id,
	})
}

// Redis类似序列化操作
func (PlayerStats PlayerStats) MarshalBinary() ([]byte, error) {
	return json.Marshal(PlayerStats)
}

func (PlayerStats PlayerStats) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &PlayerStats)
}
