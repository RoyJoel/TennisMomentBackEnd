package model

import (
	"encoding/json"

	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"
)

type Event struct {
	Id        int64            `json:"id"`
	Icon      string           `json:"icon"`
	Name      string           `json:"name"`
	StartDate float64          `json:"startDate"`
	EndDate   float64          `json:"endDate"`
	Level     int64            `json:"level"`
	Draw      utils.IntMatrix  `json:"draw"`
	Schedule  utils.IntMatrix2 `json:"schedule"`
}

func (Event Event) TableName() string {
	return "Event"
}

func (Event Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        Event.Id,
		"name":      Event.Name,
		"icon":      Event.Icon,
		"startDate": Event.StartDate,
		"endDate":   Event.EndDate,
		"level":     Event.Level,
		"draw":      Event.Draw,
		"schedule":  Event.Schedule,
	})
}

// Redis类似序列化操作
func (Event Event) MarshalBinary() ([]byte, error) {
	return json.Marshal(Event)
}

func (Event Event) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Event)
}

type EventResponse struct {
	Id        int64            `json:"id"`
	Icon      string           `json:"icon"`
	Name      string           `json:"name"`
	StartDate float64          `json:"startDate"`
	EndDate   float64          `json:"endDate"`
	Level     int64            `json:"level"`
	Draw      []PlayerResponse `json:"draw"`
	Schedule  [][]GameResponse `json:"schedule"`
}

func (EventResponse EventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        EventResponse.Id,
		"name":      EventResponse.Name,
		"icon":      EventResponse.Icon,
		"startDate": EventResponse.StartDate,
		"endDate":   EventResponse.EndDate,
		"level":     EventResponse.Level,
		"draw":      EventResponse.Draw,
		"schedule":  EventResponse.Schedule,
	})
}

// Redis类似序列化操作
func (EventResponse EventResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(EventResponse)
}

func (EventResponse EventResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &EventResponse)
}
