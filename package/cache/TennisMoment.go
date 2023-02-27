package cache

import (
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type TennisMomentCacheDAOImpl struct {
	db *redis.Client
}

type TennisMomentCacheDAO interface {
	// set一个
	SetPlayer(ctx context.Context, key string, info model.Player, t time.Duration) bool
	// 根据ID获取一个
	GetPlayerById(ctx context.Context, key string) model.Player
}

func NewTennisMomentCacheDAOImpl() *TennisMomentCacheDAOImpl {
	return &TennisMomentCacheDAOImpl{db: config.RDB}
}

func (impl TennisMomentCacheDAOImpl) SetPlayer(ctx context.Context, key string, info model.Player, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl TennisMomentCacheDAOImpl) GetPlayerById(ctx context.Context, key string) model.Player {
	res := impl.db.Get(ctx, key)
	var info model.Player
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}
