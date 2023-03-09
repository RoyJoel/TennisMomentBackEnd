package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"github.com/go-redis/redis/v8"
)

type PlayerCacheDAOImpl struct {
	db *redis.Client
}

type PlayerCacheDAO interface {
	// set一个
	SetPlayer(ctx context.Context, key string, info model.Player, t time.Duration) bool
	// 根据ID获取一个
	GetPlayerById(ctx context.Context, key string) model.Player
}

func NewPlayerCacheDAOImpl() *PlayerCacheDAOImpl {
	return &PlayerCacheDAOImpl{db: config.RDB}
}

func (impl PlayerCacheDAOImpl) SetPlayer(ctx context.Context, key string, info model.Player, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl PlayerCacheDAOImpl) GetPlayer(ctx context.Context, key string) model.Player {
	res := impl.db.Get(ctx, key)
	var info model.Player
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}
