package sms_resolver

import (
	"notification/internal/config"
	"notification/internal/constants"
	"notification/internal/domain"
	"notification/pkg/smsproviders"
	provider1 "notification/pkg/smsproviders/provider1"
	provider2 "notification/pkg/smsproviders/provider2"

	log "notification/pkg/logger"
)

type Resolver struct {
	cfg    config.SMSProvider
	appEnv config.AppEnv
	logger log.Logger
}

func NewResolver(cfg config.SMSProvider, appEnv config.AppEnv, logger log.Logger) Resolver {
	return Resolver{
		cfg:    cfg,
		appEnv: appEnv,
		logger: logger,
	}
}

func (r *Resolver) ResolveSmsProvider(provider domain.Provider, templateID string) (smsproviders.SmsSender, error) {
	var (
		err    error
		driver smsproviders.SmsSender
	)

	switch provider.Name {
	case constants.SMS_PROVIDER1:
		driver = provider1.NewProvider1()
	case constants.SMS_PROVIDER2:
		driver = provider2.NewProvider2()
	default:
		err = constants.ErrProviderNotFound
	}

	return driver, err
}
