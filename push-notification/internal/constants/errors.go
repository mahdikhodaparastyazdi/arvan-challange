package constants

import "errors"

const (
	providerNotFoundErrMsg             = "provider not found"
	providerIsNotActiveErrMsg          = "provider is not active"
	templateParamsIsRequiredErrMes     = "template params are required"
	smsProviderErrMsg                  = "internal request in provider"
	paramsAreNotProvidedErrMsg         = "params are not provided"
	templateOrProviderNotDefinedErrMsg = "provider of template not defined"
	templateNameMustBeUniqueErrMsg     = "template name is not unique"
	unsupportedSmsTypeErrMsg           = "unsupported sms type"
	templateNotFoundErrMsg             = "template not found"
	wrongApiKeyErrMsg                  = "wrong api key"
	templateProviderNotDefinedErrMsg   = "provider template not defined"
	deviceInfoNotfoundErrMsg           = "device info not found"
	userIsNotInWhiteListMsg            = "user not allowed to receive notification"
	expiryDateTimeErrMsg               = "expires at can't be for a past time"
	expiryReachedErrMsg                = "message expires"
	pushNotificationNotFoundErrMsg     = "push notification not found"
	smsNotFoundErrMsg                  = "sms not found"
	backOffRetryErrMsg                 = "backoff from retrying"
	wrongStatus                        = "wrong status"
)

var (
	ErrProviderNotFound             = errors.New(providerNotFoundErrMsg)
	ErrProviderIsNotActive          = errors.New(providerIsNotActiveErrMsg)
	ErrTemplateParamsIsRequired     = errors.New(templateParamsIsRequiredErrMes)
	ErrInternalSmsProviderError     = errors.New(smsProviderErrMsg)
	ErrParamsAreNotProvided         = errors.New(paramsAreNotProvidedErrMsg)
	ErrTemplateOrProviderNotDefined = errors.New(templateOrProviderNotDefinedErrMsg)
	ErrTemplateNameIsNotUnique      = errors.New(templateNameMustBeUniqueErrMsg)
	ErrUnsupportedSmsType           = errors.New(unsupportedSmsTypeErrMsg)
	ErrTemplateNotFound             = errors.New(templateNotFoundErrMsg)
	ErrWrongApiKey                  = errors.New(wrongApiKeyErrMsg)
	ErrTemplateProviderNotDefined   = errors.New(templateProviderNotDefinedErrMsg)
	ErrDeviceInfoNotFound           = errors.New(deviceInfoNotfoundErrMsg)
	ErrUserIsNotInWhiteList         = errors.New(userIsNotInWhiteListMsg)
	ErrExpiryDateTime               = errors.New(expiryDateTimeErrMsg)
	ErrExpiryReached                = errors.New(expiryReachedErrMsg)
	ErrPushNotificationNotfound     = errors.New(pushNotificationNotFoundErrMsg)
	ErrSmsNotFound                  = errors.New(smsNotFoundErrMsg)
	ErrBackOffRetry                 = errors.New(backOffRetryErrMsg)
	ErrWrongStatus                  = errors.New(wrongStatus)
)
