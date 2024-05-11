package repositories

import (
	"context"
	"notification/internal/domain"
	"notification/internal/model"

	"gorm.io/gorm"
)

type commonBehaviour[T model.DbModel, K domain.Domain] struct {
	db *gorm.DB
}

func NewCommonBehaviour[T model.DbModel, K domain.Domain](db *gorm.DB) CommonBehaviourRepository[T, K] {
	return &commonBehaviour[T, K]{
		db: db,
	}
}

func (c *commonBehaviour[T, K]) ByID(ctx context.Context, id uint) (K, error) {
	return c.ByField(ctx, "id", id)
}

func (c *commonBehaviour[T, K]) ByField(ctx context.Context, field string, id uint) (K, error) {
	var t T
	err := c.db.WithContext(ctx).Where(field+"=?", id).First(&t).Error

	d := t.ToDomain().(K)
	return d, err
}

func (c *commonBehaviour[T, K]) Save(ctx context.Context, model *T) error {
	return c.db.WithContext(ctx).Save(model).Error
}
