package model

import (
	"notification/internal/domain"
	"time"
)

type Template struct {
	ID                  uint      `gorm:"column:id"`
	Code                string    `gorm:"column:code"`
	Content             string    `gorm:"column:content"`
	Params              string    `gorm:"column:params"`
	UseProviderTemplate bool      `gorm:"column:use_provider_template"`
	ActiveProviderID    uint      `gorm:"column:active_provider_id"`
	Priority            string    `gorm:"column:priority"`
	Type                string    `gorm:"column:type"`
	Provider            Provider  `gorm:"foreignKey:active_provider_id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
}

func (t Template) TableName() string {
	return "templates"
}

func (t Template) ToDomain() domain.Domain {
	return domain.Template{
		ID:                  t.ID,
		Code:                t.Code,
		Content:             t.Content,
		Params:              t.Params,
		UseProviderTemplate: t.UseProviderTemplate,
		ActiveProviderID:    t.ActiveProviderID,
		Priority:            domain.TemplatePriority(t.Priority),
		Type:                domain.TemplateType(t.Type),
		CreatedAt:           t.CreatedAt,
		UpdatedAt:           t.UpdatedAt,
	}
}
