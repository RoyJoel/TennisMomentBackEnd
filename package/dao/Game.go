package dao

import (
	"context"

	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
)

type GameDAO interface {
	//添加一个
	AddGame(ctx context.Context, info model.Game) bool
	//根据Key获取一个
	SearchGame(ctx context.Context, url string) model.Game
	//查看全部
	UpdateGame(ctx context.Context, page int, limit int) model.Game
}
