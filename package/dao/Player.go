package dao

import (
	"context"

	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
)

type PlayerDAO interface {
	//添加一个
	AddPlayer(ctx context.Context, info model.Player) bool
	// 根据ID获取一个
	GetPlayerInfo(ctx context.Context, loginName string) model.Player
	// 更新球员数据
	UpdatePlayer(ctx context.Context, loginName string) model.Player
	// //根据Key获取一个
	// GetPlayerByKey(ctx context.Context, url string) model.Player
	//查看全部
	// FindAllPlayers(ctx context.Context, page int, limit int) []model.Player
	// //根据Key修改
	// UpdatePlayerByKey(ctx context.Context, info model.Player) bool
	// //删除一个
	// DeletePlayerById(ctx context.Context, id int64) bool
	//根据ID修改
	// UpdatePlayerById(ctx context.Context, info model.Player) bool
}
