package model

import "notification/internal/domain"

type ProviderTemplate struct {
	ID         uint     `gorm:"column:id"`
	ProviderID uint     `gorm:"column:provider_id"`
	TemplateID uint     `gorm:"foreignKey:Template,column:template_id"`
	Provider   Provider `gorm:"foreignKey:provider_id"`
	Template   Template `gorm:"foreignKey:template_id"`
	Code       string   `gorm:"column:code"`
}

func (p ProviderTemplate) TableName() string {
	return "provider_templates"
}

func (p ProviderTemplate) ToDomain() domain.Domain {
	return domain.ProviderTemplate{
		ID:         p.ID,
		ProviderID: p.ProviderID,
		TemplateID: p.TemplateID,
		Code:       p.Code,
	}
}
