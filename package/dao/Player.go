package dao

import (
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"context"
)

type PlayerDAO interface {
	//添加一个
	AddPlayer(ctx context.Context, info model.Player) bool
	// //根据Key获取一个
	// GetPlayerByKey(ctx context.Context, url string) model.Player
	// //查看全部
	// FindAllPlayer(ctx context.Context, page int, limit int) []model.Player
	// //根据Key修改
	// UpdatePlayerByKey(ctx context.Context, info model.Player) bool
	// //删除一个
	// DeletePlayerById(ctx context.Context, id int64) bool
	//根据ID获取一个
	GetPlayerById(ctx context.Context, id int64) model.Player
	//根据ID修改
	UpdatePlayerById(ctx context.Context, info model.Player) bool
}
