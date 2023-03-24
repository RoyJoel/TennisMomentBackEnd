package model

import (
	"encoding/json"
)

type ClubManager struct {
	ClubId    int64 `json:"clubId"`
	ManagerId int64 `json:"ManagerId"`
}

func (ClubManager ClubManager) TableName() string {
	return "ClubManager"
}

func (ClubManager ClubManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"clubId":    ClubManager.ClubId,
		"ManagerId": ClubManager.ManagerId,
	})
}

// Redis类似序列化操作
func (ClubManager ClubManager) MarshalBinary() ([]byte, error) {
	return json.Marshal(ClubManager)
}

func (ClubManager ClubManager) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &ClubManager)
}
