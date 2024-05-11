package model

import (
	"notification/internal/domain"
	"time"
)

type Provider struct {
	ID        uint      `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Type      string    `gorm:"column:type"`
	Active    bool      `gorm:"column:active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (p Provider) TableName() string {
	return "providers"
}

func (p Provider) ToDomain() domain.Domain {
	return domain.Provider{
		ID:        p.ID,
		Name:      p.Name,
		Active:    p.Active,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
