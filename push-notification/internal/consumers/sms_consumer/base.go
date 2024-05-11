package sms_consumer

import (
	"context"
	"notification/internal/config"
	"notification/internal/domain"
	"notification/internal/factories/sms_resolver"

	log "notification/pkg/logger"
)

type smsRepository interface {
	GetByID(c context.Context, id uint) (domain.SMS, error)
	UpdateStatus(c context.Context, id uint, status domain.NotificationStatus) error
}

type Consumer struct {
	smsRepository    smsRepository
	providerResolver sms_resolver.Resolver
	cfg              config.SMSProvider
	logger           log.Logger
}

func New(cfg config.SMSProvider, logger log.Logger, repository smsRepository, resolver sms_resolver.Resolver) Consumer {
	return Consumer{
		smsRepository:    repository,
		providerResolver: resolver,
		logger:           logger,
		cfg:              cfg,
	}
}
