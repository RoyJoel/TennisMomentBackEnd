package model

import (
	"encoding/json"
)

type EventPlayer struct {
	EventId  int64 `json:"eventId"`
	PlayerId int64 `json:"playerId"`
}

func (EventPlayer EventPlayer) TableName() string {
	return "EventPlayer"
}

func (EventPlayer EventPlayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"eventId":  EventPlayer.EventId,
		"playerId": EventPlayer.PlayerId,
	})
}

// Redis类似序列化操作
func (EventPlayer EventPlayer) MarshalBinary() ([]byte, error) {
	return json.Marshal(EventPlayer)
}

func (EventPlayer EventPlayer) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &EventPlayer)
}
