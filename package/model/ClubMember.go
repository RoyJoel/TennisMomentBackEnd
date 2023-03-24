package model

import (
	"encoding/json"
)

type ClubMember struct {
	ClubId   int64 `json:"clubId"`
	MemberId int64 `json:"memberId"`
}

func (ClubMember ClubMember) TableName() string {
	return "ClubMember"
}

func (ClubMember ClubMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"clubId":   ClubMember.ClubId,
		"memberId": ClubMember.MemberId,
	})
}

// Redis类似序列化操作
func (ClubMember ClubMember) MarshalBinary() ([]byte, error) {
	return json.Marshal(ClubMember)
}

func (ClubMember ClubMember) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &ClubMember)
}
