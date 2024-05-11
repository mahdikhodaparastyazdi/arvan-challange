package model

import (
	"notification/internal/domain"
	"time"
)

type SMS struct {
	ID         uint      `gorm:"column:id"`
	Mobile     string    `gorm:"column:mobile"`
	Content    string    `gorm:"column:content"`
	Status     string    `gorm:"column:status"`
	ProviderId uint      `gorm:"column:provider_id"`
	TemplateID uint      `gorm:"column:template_id"`
	Provider   Provider  `gorm:"foreignKey:provider_id"`
	Template   Template  `gorm:"foreignKey:template_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	ExpiredAt  time.Time `gorm:"column:expires_at"`
}

func (s SMS) TableName() string {
	return "sms"
}

func (s SMS) ToDomain() domain.Domain {
	return domain.SMS{
		ID:         s.ID,
		Mobile:     s.Mobile,
		Content:    s.Content,
		Status:     domain.NotificationStatus(s.Status),
		ProviderId: s.ProviderId,
		TemplateID: s.TemplateID,
		ExpiresAt:  s.ExpiredAt,
	}
}
