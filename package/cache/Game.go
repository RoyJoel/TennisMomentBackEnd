package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"github.com/go-redis/redis/v8"
)

type GameCacheDAOImpl struct {
	db *redis.Client
}

type GameCacheDAO interface {
	// set一个
	SetGame(ctx context.Context, key string, info model.Game, t time.Duration) bool
	// 根据ID获取一个
	GetGameById(ctx context.Context, key string) model.Game
}

func NewGameCacheDAOImpl() *GameCacheDAOImpl {
	return &GameCacheDAOImpl{db: config.RDB}
}

func (impl GameCacheDAOImpl) SetGame(ctx context.Context, key string, info model.Game, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl GameCacheDAOImpl) GetGame(ctx context.Context, key string) model.Game {
	res := impl.db.Get(ctx, key)
	var info model.Game
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}
