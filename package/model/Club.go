package model

import (
	"encoding/json"

	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"
)

type Club struct {
	Id      int64           `json:"id"`
	Name    string          `json:"name"`
	Icon    string          `json:"icon"`
	Intro   string          `json:"intro"`
	OwnerId int64           `json:"ownerId"`
	Address string          `json:"address"`
	Events  utils.IntMatrix `json:"events"`
}

func (Club Club) TableName() string {
	return "Club"
}

func (Club Club) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":      Club.Id,
		"name":    Club.Name,
		"icon":    Club.Icon,
		"intro":   Club.Intro,
		"ownerId": Club.OwnerId,
		"address": Club.Address,
	})
}

// Redis类似序列化操作
func (Club Club) MarshalBinary() ([]byte, error) {
	return json.Marshal(Club)
}

func (Club Club) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Club)
}

type ClubResponse struct {
	Id      int64                `json:"id"`
	Name    string               `json:" name"`
	Icon    string               `json:"icon"`
	Intro   string               `json:"events"`
	Owner   PlayerResponse `json:"ownerId"`
	Address string               `json:"address"`
	Events  []EventResponse      `json:"events"`
}

func (ClubResponse ClubResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":      ClubResponse.Id,
		"name":    ClubResponse.Name,
		"icon":    ClubResponse.Icon,
		"intro":   ClubResponse.Intro,
		"owner":   ClubResponse.Owner,
		"address": ClubResponse.Address,
		"events":  ClubResponse.Events,
	})
}

// Redis类似序列化操作
func (ClubResponse ClubResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(ClubResponse)
}

func (ClubResponse ClubResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &ClubResponse)
}
