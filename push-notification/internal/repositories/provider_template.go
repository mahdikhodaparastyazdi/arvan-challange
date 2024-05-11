package repositories

import (
	"context"
	"errors"
	"notification/internal/constants"
	"notification/internal/domain"
	"notification/internal/model"

	"gorm.io/gorm"
)

type providerTemplateRepository struct {
	CommonBehaviourRepository[model.ProviderTemplate, domain.ProviderTemplate]
	db *gorm.DB
}

func NewProviderTemplate(db *gorm.DB) ProviderTemplateRepository {
	return providerTemplateRepository{
		CommonBehaviourRepository: NewCommonBehaviour[model.ProviderTemplate, domain.ProviderTemplate](db),
		db:                        db,
	}
}

func (p providerTemplateRepository) Create(ctx context.Context, pt domain.ProviderTemplate) (domain.ProviderTemplate, error) {
	mPt := model.ProviderTemplate{
		ProviderID: pt.ProviderID,
		TemplateID: pt.TemplateID,
		Code:       pt.Code,
	}

	if err := p.Save(ctx, &mPt); err != nil {
		return domain.ProviderTemplate{}, err
	}

	dMpt := mPt.ToDomain().(domain.ProviderTemplate)

	return dMpt, nil
}

func (p providerTemplateRepository) GetByID(ctx context.Context, id uint) (domain.ProviderTemplate, error) {
	return p.ByID(ctx, id)
}

func (p providerTemplateRepository) GetProviderTemplate(ctx context.Context, providerID, templateID uint) (domain.ProviderTemplate, error) {
	var providerTmpl model.ProviderTemplate
	err := p.db.WithContext(ctx).Where("provider_id = ? and template_id = ?", providerID, templateID).First(&providerTmpl).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ProviderTemplate{}, constants.ErrTemplateProviderNotDefined
		}
		return domain.ProviderTemplate{}, err
	}
	dt := providerTmpl.ToDomain().(domain.ProviderTemplate)

	return dt, nil
}

func (p providerTemplateRepository) GetAll(ctx context.Context) ([]domain.ProviderTemplate, error) {
	pt := make([]model.ProviderTemplate, 0)

	err := p.db.WithContext(ctx).Model(&model.ProviderTemplate{}).Find(&pt).Error
	if err != nil {
		return nil, err
	}

	dProviders := make([]domain.ProviderTemplate, 0)
	for _, v := range pt {
		vl := v.ToDomain().(domain.ProviderTemplate)
		dProviders = append(dProviders, vl)
	}

	return dProviders, nil
}
