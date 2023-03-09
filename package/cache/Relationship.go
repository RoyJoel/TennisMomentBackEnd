package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"github.com/go-redis/redis/v8"
)

type RelationshipCacheDAOImpl struct {
	db *redis.Client
}

type RelationshipCacheDAO interface {
	// set一个
	SetRelationship(ctx context.Context, key string, info model.Relationship, t time.Duration) bool
	// 根据ID获取一个
	GetRelationshipById(ctx context.Context, key string) model.Relationship
}

func NewRelationshipCacheDAOImpl() *RelationshipCacheDAOImpl {
	return &RelationshipCacheDAOImpl{db: config.RDB}
}

func (impl RelationshipCacheDAOImpl) SetRelationship(ctx context.Context, key string, info model.Relationship, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl RelationshipCacheDAOImpl) GetRelationship(ctx context.Context, key string) model.Relationship {
	res := impl.db.Get(ctx, key)
	var info model.Relationship
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}
