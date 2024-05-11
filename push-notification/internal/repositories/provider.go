package repositories

import (
	"context"
	"errors"
	"notification/internal/constants"
	"notification/internal/domain"
	"notification/internal/model"

	"gorm.io/gorm"
)

type providerRepository struct {
	CommonBehaviourRepository[model.Provider, domain.Provider]
	db *gorm.DB
}

func NewProviderRepository(db *gorm.DB) ProviderRepository {
	return providerRepository{
		CommonBehaviourRepository: NewCommonBehaviour[model.Provider, domain.Provider](db),
		db:                        db,
	}
}

func (p providerRepository) GetByID(ctx context.Context, id uint) (domain.Provider, error) {
	return p.ByID(ctx, id)
}

func (p providerRepository) ToggleActivationStatus(ctx context.Context, id uint, status bool) error {
	return p.db.WithContext(ctx).Model(&model.Provider{}).Where("id = ?", id).Update("active", status).Error
}

func (p providerRepository) GetAll(ctx context.Context) ([]domain.Provider, error) {
	providers := make([]model.Provider, 0)

	err := p.db.WithContext(ctx).Model(&model.Provider{}).Find(&providers).Error
	if err != nil {
		return nil, err
	}

	dProviders := make([]domain.Provider, 0)
	for _, v := range providers {
		vl := v.ToDomain().(domain.Provider)
		dProviders = append(dProviders, vl)
	}

	return dProviders, nil
}

func (p providerRepository) GetActiveProviderByType(ctx context.Context, providerType domain.ProviderType) (domain.Provider, error) {
	var pr model.Provider
	err := p.db.WithContext(ctx).Where("type = ?", providerType).Where("active = ?", true).First(&pr).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Provider{}, constants.ErrProviderNotFound
		}
		return domain.Provider{}, err
	}

	return pr.ToDomain().(domain.Provider), nil
}
