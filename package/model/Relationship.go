package model

import (
	"encoding/json"
)

type Relationship struct {
	LoginName       string `json:"loginName"`
	FriendLoginName string `json:"FriendLoginName"`
}

func (relationship Relationship) TableName() string {
	return "Relationship"
}

func (relationship Relationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"loginName":       relationship.LoginName,
		"friendLoginName": relationship.FriendLoginName,
	})
}

// Redis类似序列化操作
func (relationship Relationship) MarshalBinary() ([]byte, error) {
	return json.Marshal(relationship)
}

func (relationship Relationship) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &relationship)
}
