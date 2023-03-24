package model

import (
	"encoding/json"
)

type EventGame struct {
	EventId int64 `json:"eventId"`
	GameId  int64 `json:"gameId"`
}

func (EventGame EventGame) TableName() string {
	return "EventGame"
}

func (EventGame EventGame) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"eventId": EventGame.EventId,
		"gameId":  EventGame.GameId,
	})
}

// Redis类似序列化操作
func (EventGame EventGame) MarshalBinary() ([]byte, error) {
	return json.Marshal(EventGame)
}

func (EventGame EventGame) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &EventGame)
}
