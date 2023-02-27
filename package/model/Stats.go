package model

import (
	"encoding/json"
)

type Stats struct {
	Aces                       int64 `json:"aces"`
	DoubleFaults               int64 `json:"doubleFaults"`
	FirstServePoints           int64 `json:"firstServePoints"`
	FirstServePointsIn         int64 `json:"firstServePointsIn"`
	FirstServePointsWon        int64 `json:"firstServePointsWon"`
	SecondServePoints          int64 `json:"secondServePoints"`
	SecondServePointsWon       int64 `json:"secondServePointsWo"`
	BreakPointsFaced           int64 `json:"breakPointsFaced"`
	BreakPointsSaved           int64 `json:"breakPointsSaved"`
	ServeGamesPlayed           int64 `json:"serveGamesPlayed"`
	ServeGamesWon              int64 `json:"serveGamesWon"`
	TotalServePointsWon        int64 `json:"totalServePointsWon"`
	FirstServeReturnPoints     int64 `json:"firstServeReturnPoints"`
	FirstServeReturnAces       int64 `json:"firstServeReturnAces"`
	FirstServeReturnPointsWon  int64 `json:"firstServeReturnPointsWon"`
	SecondServeReturnPoints    int64 `json:"secondServeReturnPoints"`
	SecondServeReturnAces      int64 `json:"secondServeReturnAces"`
	SecondServeReturnPointsWon int64 `json:"secondServeReturnPointsWon"`
	BreakPointsOpportunities   int64 `json:"breakPointsOpportunities"`
	BreakPointsConverted       int64 `json:"breakPointsConverted"`
	ReturnGamesPlayed          int64 `json:"returnGamesPlayed"`
	ReturnGamesWon             int64 `json:"returnGamesWon"`
	TotalReturnPointsWon       int64 `json:"totalReturnPointsWon"`
	TotalPointsWon             int64 `json:"totalPointsWon"`
	NetPoints                  int64 `json:"netPoints"`
	UnforcedErrors             int64 `json:"unforcedErrors"`
	ForehandWinners            int64 `json:"forehandWinners"`
	BackhandWinners            int64 `json:"backhandWinners"`
	PlayerId                   int64 `json:"player_id" gorm:"not null"`
	Player                     Player 
}

func (stats Stats) TableName() string {
	return "Stats"
}

func (stats Stats) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"aces":                       stats.Aces,
		"doubleFaults":               stats.DoubleFaults,
		"firstServePoints":           stats.FirstServePoints,
		"firstServePointsIn":         stats.FirstServePointsIn,
		"firstServePointsWon":        stats.FirstServePointsWon,
		"secondServePoints":          stats.SecondServePoints,
		"secondServePointsWon":       stats.SecondServePointsWon,
		"breakPointsFaced":           stats.BreakPointsFaced,
		"breakPointsSaved":           stats.BreakPointsSaved,
		"serveGamesPlayed":           stats.ServeGamesPlayed,
		"serveGamesWon":              stats.ServeGamesWon,
		"totalServePointsWon":        stats.TotalServePointsWon,
		"firstServeReturnPoints":     stats.FirstServeReturnPoints,
		"firstServeReturnAces":       stats.FirstServeReturnAces,
		"firstServeReturnPointsWon":  stats.FirstServeReturnPointsWon,
		"secondServeReturnPoints":    stats.SecondServeReturnPoints,
		"secondServeReturnAces":      stats.SecondServeReturnAces,
		"secondServeReturnPointsWon": stats.SecondServeReturnPointsWon,
		"breakPointsOpportunities":   stats.BreakPointsOpportunities,
		"breakPointsConverted":       stats.BreakPointsConverted,
		"returnGamesPlayed":          stats.ReturnGamesPlayed,
		"returnGamesWon":             stats.ReturnGamesWon,
		"totalReturnPointsWon":       stats.TotalReturnPointsWon,
		"totalPointsWon":             stats.TotalPointsWon,
		"netPoints":                  stats.NetPoints,
		"unforcedErrors":             stats.UnforcedErrors,
		"forehandWinners":            stats.ForehandWinners,
		"backhandWinners":            stats.BackhandWinners,
		"player_id":                  stats.PlayerId,
	})
}

// Redis类似序列化操作
func (stats Stats) MarshalBinary() ([]byte, error) {
	return json.Marshal(stats)
}

func (stats Stats) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &stats)
}
