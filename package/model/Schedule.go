package model

import (
	"encoding/json"
)

type Schedule struct {
	Id        int64   `json:"id"`
	Place     string  `json:"place"`
	StartDate float64 `json:"startDate"`
	Player1Id int64   `json:"player1Id"`
	Player2Id int64   `json:"player2Id"`
}

func (Schedule Schedule) TableName() string {
	return "Schedule"
}

func (Schedule Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        Schedule.Id,
		"place":     Schedule.Place,
		"startDate": Schedule.StartDate,
		"player1Id": Schedule.Player1Id,
		"player2Id": Schedule.Player2Id,
	})
}

// Redis类似序列化操作
func (Schedule Schedule) MarshalBinary() ([]byte, error) {
	return json.Marshal(Schedule)
}

func (Schedule Schedule) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Schedule)
}

type ScheduleResponse struct {
	Id        int64          `json:"id"`
	Place     string         `json:"place"`
	StartDate float64        `json:"startDate"`
	Opponent  PlayerResponse `json:"opponent"`
}

func (scheduleResponse ScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        scheduleResponse.Id,
		"place":     scheduleResponse.Place,
		"startDate": scheduleResponse.StartDate,
		"opponent":  scheduleResponse.Opponent,
	})
}

// Redis类似序列化操作
func (scheduleResponse ScheduleResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(scheduleResponse)
}

func (scheduleResponse ScheduleResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &scheduleResponse)
}
