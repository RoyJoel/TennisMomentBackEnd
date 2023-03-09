package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"github.com/go-redis/redis/v8"
)

type StatsCacheDAOImpl struct {
	db *redis.Client
}

type StatsCacheDAO interface {
	// set一个
	SetStats(ctx context.Context, key string, info model.Stats, t time.Duration) bool
	// 根据ID获取一个
	GetStatsById(ctx context.Context, key string) model.Stats
}

func NewStatsCacheDAOImpl() *StatsCacheDAOImpl {
	return &StatsCacheDAOImpl{db: config.RDB}
}

func (impl StatsCacheDAOImpl) SetStats(ctx context.Context, key string, info model.Stats, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl StatsCacheDAOImpl) GetStats(ctx context.Context, key string) model.Stats {
	res := impl.db.Get(ctx, key)
	var info model.Stats
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}
