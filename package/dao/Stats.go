package dao

import (
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"context"
)

type StatsDAO interface {
	//添加一个
	AddStats(ctx context.Context, info model.Stats) bool
	// //根据Key获取一个
	// GetStatsByKey(ctx context.Context, url string) model.Stats
	// //查看全部
	// FindAllStats(ctx context.Context, page int, limit int) []model.Stats
	// //根据Key修改
	// UpdateStatsByKey(ctx context.Context, info model.Stats) bool
	// //删除一个
	// DeleteStatsById(ctx context.Context, id int64) bool
	//根据ID获取一个
	GetStatsById(ctx context.Context, id int64) model.Stats
	//根据ID修改
	UpdateStatsById(ctx context.Context, info model.Stats) bool
}
