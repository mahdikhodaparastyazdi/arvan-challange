package domain

import "time"

type SMS struct {
	ID         uint
	ProviderId uint
	TemplateID uint
	Mobile     string
	Content    string
	Status     NotificationStatus
	ExpiresAt  time.Time
}

func (s SMS) IsDomain() bool {
	return true
}
