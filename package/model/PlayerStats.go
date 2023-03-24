package model

import (
	"encoding/json"
)

type PlayerStats struct {
	PlayerId int64 `json:"playerId"`
	StatsId        int64  `json:"statsId"`
}

func (PlayerStats PlayerStats) TableName() string {
	return "PlayerStats"
}

func (PlayerStats PlayerStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"playerId": PlayerStats.PlayerId,
		"statsId":        PlayerStats.StatsId,
	})
}

// Redis类似序列化操作
func (PlayerStats PlayerStats) MarshalBinary() ([]byte, error) {
	return json.Marshal(PlayerStats)
}

func (PlayerStats PlayerStats) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &PlayerStats)
}
