package repositories

import (
	"context"
	"errors"
	"notification/internal/constants"
	"notification/internal/domain"
	"notification/internal/model"

	"gorm.io/gorm"
)

type templateRepository struct {
	CommonBehaviourRepository[model.Template, domain.Template]
	db *gorm.DB
}

func NewTemplateRepository(db *gorm.DB) TemplateRepository {
	return templateRepository{
		CommonBehaviourRepository: NewCommonBehaviour[model.Template, domain.Template](db),
		db:                        db,
	}
}

func (t templateRepository) Create(ctx context.Context, tmpl domain.Template) (domain.Template, error) {
	tm := model.Template{
		Code:                tmpl.Code,
		Content:             tmpl.Content,
		Params:              tmpl.Params,
		UseProviderTemplate: tmpl.UseProviderTemplate,
		ActiveProviderID:    tmpl.ActiveProviderID,
		Priority:            string(tmpl.Priority),
		Type:                string(tmpl.Type),
	}
	if err := t.db.WithContext(ctx).Model(&model.Template{}).Create(&tm).Error; err != nil {
		return tmpl, err
	}

	return tmpl, nil
}
func (t templateRepository) GetByID(ctx context.Context, id uint) (domain.Template, error) {
	return t.ByID(ctx, id)
}

func (t templateRepository) FindByCode(ctx context.Context, code string) (domain.Template, error) {
	var template model.Template
	err := t.db.WithContext(ctx).Where("code = ?", code).First(&template).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Template{}, constants.ErrTemplateNotFound
		}
		return domain.Template{}, err
	}

	dt := template.ToDomain().(domain.Template)

	return dt, nil
}

func (t templateRepository) Update(ctx context.Context, tmpl domain.Template) error {
	template := model.Template{
		ID:                  tmpl.ID,
		Code:                tmpl.Code,
		Content:             tmpl.Content,
		Params:              tmpl.Params,
		UseProviderTemplate: tmpl.UseProviderTemplate,
		ActiveProviderID:    tmpl.ActiveProviderID,
		Priority:            string(tmpl.Priority),
		Type:                string(tmpl.Type),
	}

	if err := t.db.WithContext(ctx).Model(&template).Updates(template).Error; err != nil {
		return err
	}

	return nil
}

func (t templateRepository) GetAll(ctx context.Context) ([]domain.Template, error) {
	mt := make([]model.Template, 0)

	err := t.db.WithContext(ctx).Model(&model.Template{}).Find(&mt).Error
	if err != nil {
		return nil, err
	}

	dTemplates := make([]domain.Template, 0)
	for _, v := range mt {
		vl := v.ToDomain().(domain.Template)
		dTemplates = append(dTemplates, vl)
	}

	return dTemplates, nil
}
