package model

import (
	"encoding/json"
)

type ClubEvent struct {
	ClubId  int64 `json:"clubId"`
	EventId int64 `json:"eventId"`
}

func (ClubEvent ClubEvent) TableName() string {
	return "ClubEvent"
}

func (ClubEvent ClubEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"clubId":  ClubEvent.ClubId,
		"eventId": ClubEvent.EventId,
	})
}

// Redis类似序列化操作
func (ClubEvent ClubEvent) MarshalBinary() ([]byte, error) {
	return json.Marshal(ClubEvent)
}

func (ClubEvent ClubEvent) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &ClubEvent)
}
