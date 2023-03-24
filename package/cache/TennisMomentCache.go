package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"github.com/go-redis/redis/v8"
)

type TennisMomentCacheDAOImpl struct {
	db *redis.Client
}

type TennisMomentCacheDAO interface {
	// set一个
	SetGame(ctx context.Context, key string, info model.Game, t time.Duration) bool
	// 根据ID获取一个
	GetGameById(ctx context.Context, key string) model.Game
}

func NewTennisMomentCacheDAOImpl() *TennisMomentCacheDAOImpl {
	return &TennisMomentCacheDAOImpl{db: config.RDB}
}

func (impl TennisMomentCacheDAOImpl) SetGame(ctx context.Context, key string, info model.Game, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl TennisMomentCacheDAOImpl) GetGame(ctx context.Context, key string) model.Game {
	res := impl.db.Get(ctx, key)
	var info model.Game
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}

func (impl TennisMomentCacheDAOImpl) SetPlayer(ctx context.Context, key string, info model.Player, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl TennisMomentCacheDAOImpl) GetPlayer(ctx context.Context, key string) model.Player {
	res := impl.db.Get(ctx, key)
	var info model.Player
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}

func (impl TennisMomentCacheDAOImpl) SetPlayerStats(ctx context.Context, key string, info model.PlayerStats, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl TennisMomentCacheDAOImpl) GetPlayerStats(ctx context.Context, key string) model.PlayerStats {
	res := impl.db.Get(ctx, key)
	var info model.PlayerStats
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}

func (impl TennisMomentCacheDAOImpl) SetRelationship(ctx context.Context, key string, info model.Relationship, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl TennisMomentCacheDAOImpl) GetRelationship(ctx context.Context, key string) model.Relationship {
	res := impl.db.Get(ctx, key)
	var info model.Relationship
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}

func (impl TennisMomentCacheDAOImpl) SetStats(ctx context.Context, key string, info model.Stats, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl TennisMomentCacheDAOImpl) GetStats(ctx context.Context, key string) model.Stats {
	res := impl.db.Get(ctx, key)
	var info model.Stats
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}
