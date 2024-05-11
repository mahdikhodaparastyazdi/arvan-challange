package domain

import "time"

type ProviderType string

const (
	ProviderTypeSms              ProviderType = "SMS"
	ProviderTypePushNotification ProviderType = "PUSH"
)

type Provider struct {
	ID        uint
	Name      string
	Type      ProviderType
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Provider) IsDomain() bool {
	return true
}
