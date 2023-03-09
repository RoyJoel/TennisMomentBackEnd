package model

import (
	"encoding/json"
)

type Stats struct {
	Id                         int64 `json:"id"`
	Aces                       int64 `json:"aces"`
	DoubleFaults               int64 `json:”doubleFaults"`
	ServePoints                int64 `json:"servePoints"`
	FirstServePoints           int64 `json:”firstServePoints"`
	FirstServePointsIn         int64 `json:”firstServePointsIn”`
	FirstServePointsWon        int64 `json:”firstServePointsWon"`
	SecondServePointsWon       int64 `json:”secondServePointsWo"`
	BreakPointsFaced           int64 `json:"breakPointsFaced"`
	BreakPointsSaved           int64 `json:"breakPointsSaved"`
	ServeGamesPlayed           int64 `json:"serveGamesPlayed"`
	ServeGamesWon              int64 `json:”serveGamesWon"`
	ReturnAces                 int64 `json:"returnAces"`
	ReturnServePoints          int64 `json:"returnServePoints"`
	FirstServeReturnPoints     int64 `json:"firstServeReturnPoints"`
	FirstServeReturnPointsWon  int64 `json:"firstServeReturnPointsWon"`
	SecondServeReturnPointsWon int64 `json:"secondServeReturnPointsWon"`
	BreakPointsOpportunities   int64 `json:"breakPointsOpportunities"`
	BreakPointsConverted       int64 `json:"breakPointsConverted"`
	ReturnGamesPlayed          int64 `json:"returnGamesPlayed"`
	ReturnGamesWon             int64 `json:”returnGamesWon"`
	NetPoints                  int64 `json:"netPoints"`
	UnforcedErrors             int64 `json:"unforcedErrors"`
	ForehandWinners            int64 `json:"forehandWinners"`
	BackhandWinners            int64 `json:"backhandWinners"`
}

func (stats Stats) TableName() string {
	return "Stats"
}

func (stats Stats) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                         stats.Id,
		"aces":                       stats.Aces,
		"doubleFaults":               stats.DoubleFaults,
		"servePoints":                stats.ServePoints,
		"firstServePoints":           stats.FirstServePoints,
		"firstServePointsIn":         stats.FirstServePointsIn,
		"firstServePointsWon":        stats.FirstServePointsWon,
		"secondServePointsWon":       stats.SecondServePointsWon,
		"breakPointsFaced":           stats.BreakPointsFaced,
		"breakPointsSaved":           stats.BreakPointsSaved,
		"serveGamesPlayed":           stats.ServeGamesPlayed,
		"serveGamesWon":              stats.ServeGamesWon,
		"returnAces":                 stats.ReturnAces,
		"returnServePoints":          stats.ReturnServePoints,
		"firstServeReturnPoints":     stats.FirstServeReturnPoints,
		"firstServeReturnPointsWon":  stats.FirstServeReturnPointsWon,
		"secondServeReturnPointsWon": stats.SecondServeReturnPointsWon,
		"breakPointsOpportunities":   stats.BreakPointsOpportunities,
		"breakPointsConverted":       stats.BreakPointsConverted,
		"returnGamesPlayed":          stats.ReturnGamesPlayed,
		"returnGamesWon":             stats.ReturnGamesWon,
		"netPoints":                  stats.NetPoints,
		"unforcedErrors":             stats.UnforcedErrors,
		"forehandWinners":            stats.ForehandWinners,
		"backhandWinners":            stats.BackhandWinners,
	})
}

// Redis类似序列化操作
func (stats Stats) MarshalBinary() ([]byte, error) {
	return json.Marshal(stats)
}

func (stats Stats) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &stats)
}
