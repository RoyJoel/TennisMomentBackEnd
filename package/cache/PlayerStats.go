package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"github.com/go-redis/redis/v8"
)

type PlayerStatsCacheDAOImpl struct {
	db *redis.Client
}

type PlayerStatsCacheDAO interface {
	// set一个
	SetPlayerStats(ctx context.Context, key string, info model.PlayerStats, t time.Duration) bool
	// 根据ID获取一个
	GetPlayerStatsById(ctx context.Context, key string) model.PlayerStats
}

func NewPlayerStatsCacheDAOImpl() *PlayerStatsCacheDAOImpl {
	return &PlayerStatsCacheDAOImpl{db: config.RDB}
}

func (impl PlayerStatsCacheDAOImpl) SetPlayerStats(ctx context.Context, key string, info model.PlayerStats, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl PlayerStatsCacheDAOImpl) GetPlayerStats(ctx context.Context, key string) model.PlayerStats {
	res := impl.db.Get(ctx, key)
	var info model.PlayerStats
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}
