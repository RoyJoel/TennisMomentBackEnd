package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/RoyJoel/TennisMomentBackEnd/package/cache"
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"

	"gorm.io/gorm"
)

type RelationshipDaoImpl struct {
	db    *gorm.DB
	cache *cache.RelationshipCacheDAOImpl
}

func NewRelationshipDaoImpl() *RelationshipDaoImpl {
	return &RelationshipDaoImpl{db: config.DB, cache: cache.NewRelationshipCacheDAOImpl()}
}

func (impl *RelationshipDaoImpl) AddRelationship(ctx context.Context, Relationship model.Relationship) bool {
	impl.db.Save(&Relationship)
	return true
}

// func (impl *RelationshipDaoImpl) AddRelationship(ctx context.Context, Relationship model.Relationship) bool {
// 	var p model.Relationship
// 	impl.db.First(&p, "loginName", Relationship.LoginName)
// 	if p.LoginName == Relationship.LoginName {
// 		return false
// 	}

// 	impl.db.Save(&Relationship)
// 	return true
// }

func (impl *RelationshipDaoImpl) UpdateRelationship(ctx context.Context, Relationship model.Relationship) bool {
	result := impl.db.Model(&model.Relationship{}).Where("login_name = ? ", Relationship.LoginName).Updates(Relationship)
	return result.RowsAffected > 0
}

func (impl *RelationshipDaoImpl) SearchRelationship(ctx context.Context, name string) (*model.Relationship, error) {
	var Relationship model.Relationship
	result := impl.db.First(&Relationship, "login_name = ", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Relationship not found for players: %s", name)
		}
		return nil, fmt.Errorf("failed to get Relationship for players %s: %v", name, result.Error)
	}
	return &Relationship, nil
}
