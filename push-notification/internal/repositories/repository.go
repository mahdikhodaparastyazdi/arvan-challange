package repositories

import (
	"context"
	"notification/internal/domain"
	"notification/internal/model"
	"time"
)

type CommonBehaviourRepository[T model.DbModel, K domain.Domain] interface {
	ByID(ctx context.Context, id uint) (K, error)
	ByField(ctx context.Context, field string, id uint) (K, error)
	Save(ctx context.Context, model *T) error
}

type ProviderRepository interface {
	CommonBehaviourRepository[model.Provider, domain.Provider]
	ToggleActivationStatus(ctx context.Context, id uint, status bool) error
	GetAll(ctx context.Context) ([]domain.Provider, error)
	GetByID(ctx context.Context, id uint) (domain.Provider, error)
	GetActiveProviderByType(ctx context.Context, providerType domain.ProviderType) (domain.Provider, error)
}

type ProviderTemplateRepository interface {
	CommonBehaviourRepository[model.ProviderTemplate, domain.ProviderTemplate]
	Create(ctx context.Context, pt domain.ProviderTemplate) (domain.ProviderTemplate, error)
	GetProviderTemplate(c context.Context, providerID, templateID uint) (domain.ProviderTemplate, error)
	GetByID(ctx context.Context, id uint) (domain.ProviderTemplate, error)
	GetAll(ctx context.Context) ([]domain.ProviderTemplate, error)
}

type SmsRepository interface {
	CommonBehaviourRepository[model.SMS, domain.SMS]
	UpdateStatus(c context.Context, id uint, status domain.NotificationStatus) error
	Create(ctx context.Context, sms domain.SMS) (domain.SMS, error)
	GetByID(ctx context.Context, id uint) (domain.SMS, error)
	CountSince(ctx context.Context, mobile string, since time.Time) (int, error)
}

type TemplateRepository interface {
	CommonBehaviourRepository[model.Template, domain.Template]
	FindByCode(ctx context.Context, code string) (domain.Template, error)
	Create(ctx context.Context, tmpl domain.Template) (domain.Template, error)
	Update(ctx context.Context, tmpl domain.Template) error
	GetAll(ctx context.Context) ([]domain.Template, error)
	GetByID(ctx context.Context, id uint) (domain.Template, error)
}

type ClientRepository interface {
	CommonBehaviourRepository[model.Client, domain.Client]
	GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error)
	UpdateBalance(ctx context.Context, amount int, clientId uint) error
}
