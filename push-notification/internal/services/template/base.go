package template

import (
	"context"
	"notification/internal/config"
	"notification/internal/domain"
	"notification/internal/dto"

	log "notification/pkg/logger"
)

type Queue interface {
	Enqueue(msg dto.Message, typeName string) error
}

type smsRepository interface {
	Create(c context.Context, sms domain.SMS) (domain.SMS, error)
}
type clientRepository interface {
	UpdateBalance(ctx context.Context, amount int, clientId uint) error
}

type templateRepository interface {
	FindByCode(c context.Context, code string) (domain.Template, error)
	GetAll(ctx context.Context) ([]domain.Template, error)
	Update(ctx context.Context, tmpl domain.Template) error
	Create(ctx context.Context, tmpl domain.Template) (domain.Template, error)
}

type providerRepository interface {
	ByID(ctx context.Context, id uint) (domain.Provider, error)
}

type providerTemplateRepository interface {
	GetProviderTemplate(c context.Context, providerID, templateID uint) (domain.ProviderTemplate, error)
}

type Service struct {
	ClientRepository           clientRepository
	SmsRepository              smsRepository
	TemplateRepository         templateRepository
	providerRepository         providerRepository
	providerTemplateRepository providerTemplateRepository
	queue                      Queue
	messagesTTL                config.MessagesTTL

	logger log.Logger
}

func New(
	cr clientRepository,
	sr smsRepository,
	tr templateRepository,
	redisQueue Queue,
	pr providerRepository,
	providerTemplateRepo providerTemplateRepository,
	msgTtl config.MessagesTTL,
	logger log.Logger,
) Service {
	return Service{
		ClientRepository:           cr,
		SmsRepository:              sr,
		TemplateRepository:         tr,
		providerRepository:         pr,
		queue:                      redisQueue,
		providerTemplateRepository: providerTemplateRepo,
		messagesTTL:                msgTtl,
		logger:                     logger,
	}
}
