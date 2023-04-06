package model

import (
	"encoding/json"
)

type User struct {
	Id                int64              `json:"id"`
	LoginName         string             `json:"loginName"`
	Password          string             `json:"password"`
	Name              string             `json:"name"`
	Icon              string             `json:"icon"`
	Sex               string             `json:"sex"`
	Age               int64              `json:"age"`
	YearsPlayed       int64              `json:"yearsPlayed"`
	Height            float32            `json:"height"`
	Width             float32            `json:"width"`
	Grip              string             `json:"grip"`
	Backhand          string             `json:"backhand"`
	Points            int64              `json:"points"`
	IsAdult           bool               `json:"isAdult"`
	CareerStats       Stats              `json:"careerStats"`
	Friends           []PlayerResponse   `json:"friends"`
	AllClubs          []ClubResponse     `json:"allClubs"`
	AllGames          []GameResponse     `json:"allGames"`
	AllTourLevelGames []GameResponse     `json:"allTourLevelGames"`
	AllEvents         []EventResponse    `json:"allEvents"`
	AllSchedules      []ScheduleResponse `json:"allSchedules"`
	Token             string             `json:"token"`
}

func (User User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                User.Id,
		"loginName":         User.LoginName,
		"password":          User.Password,
		"name":              User.Name,
		"icon":              User.Icon,
		"sex":               User.Sex,
		"age":               User.Age,
		"yearsPlayed":       User.YearsPlayed,
		"height":            User.Height,
		"width":             User.Width,
		"grip":              User.Grip,
		"backhand":          User.Backhand,
		"points":            User.Points,
		"isAdult":           User.IsAdult,
		"careerStats":       User.CareerStats,
		"friends":           User.Friends,
		"allClubs":          User.AllClubs,
		"allGames":          User.AllGames,
		"allTourLevelGames": User.AllTourLevelGames,
		"allEvents":         User.AllEvents,
		"allSchedules":      User.AllSchedules,
		"token":             User.Token,
	})
}

// Redis类似序列化操作
func (User User) MarshalBinary() ([]byte, error) {
	return json.Marshal(User)
}

func (User User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &User)
}
