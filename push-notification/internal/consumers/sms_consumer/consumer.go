package sms_consumer

import (
	"context"
	"errors"
	"fmt"
	"notification/internal/constants"
	"notification/internal/domain"
	"notification/internal/dto"
	"notification/pkg/smsproviders"
	"strconv"
	"time"

	log "notification/pkg/logger"
)

type send func(context.Context, domain.SMS, dto.Message, smsproviders.SmsSender) error

func (c Consumer) Consume(ctx context.Context, message dto.Message, retry, maxRetry int) error {
	sms, err := c.smsRepository.GetByID(ctx, message.Id)
	if err != nil {
		c.logger.Error("failed to get sms", log.J{
			"providerName": message.ProviderName,
			"error":        err.Error(),
			"id":           message.Id,
		})
		if errors.Is(err, constants.ErrSmsNotFound) {
			return nil
		}
		return err
	}

	if sms.Status == domain.NotificationStatusExpired ||
		sms.Status == domain.NotificationStatusFailed ||
		sms.Status == domain.NotificationStatusSent {
		c.logger.Warn("skip send sms because of wrong state of status", log.J{
			"id":     sms.ID,
			"status": sms.Status.String(),
		})
		return constants.ErrWrongStatus
	}
	expired, err := c.isMessageExpired(ctx, sms)
	if err != nil {
		return err
	}
	if expired {
		return constants.ErrExpiryReached
	}

	driver, err := c.providerResolver.ResolveSmsProvider(domain.Provider{Name: message.ProviderName}, message.ProviderTemplateCode)
	if err != nil {
		c.logger.Error("failed to resolve sms provider", log.J{
			"providerName": message.ProviderName,
			"error":        err.Error(),
		})
		return err
	}

	var sendFunc send
	switch message.SmsType {
	case constants.SMS_TYPE_OTP:
		sendFunc = c.sendOtp
	case constants.SMS_TYPE_SMS:
		sendFunc = c.sendSms
	default:
		return constants.ErrUnsupportedSmsType
	}

	err = sendFunc(ctx, sms, message, driver)
	if err == nil {
		if err := c.smsRepository.UpdateStatus(ctx, sms.ID, domain.NotificationStatusSent); err != nil {
			c.logger.Error("failed to update sms status", log.J{
				"error": err.Error(),
				"id":    sms.ID,
			})
			return err
		}
		c.logger.Info("sms send", log.J{
			"id": sms.ID,
		})
		return nil
	}

	var status domain.NotificationStatus
	if retry >= 0 && retry < maxRetry {
		status = domain.NotificationStatusRetry
	} else {
		status = domain.NotificationStatusFailed
	}

	if sms.Status != status {
		if err := c.smsRepository.UpdateStatus(ctx, sms.ID, status); err != nil {
			c.logger.Error("failed to update sms status", log.J{
				"error": err.Error(),
				"id":    sms.ID,
			})
			return err
		}
	}
	if status == domain.NotificationStatusFailed {
		c.logger.Warn("backoff from retrying sms notification", log.J{
			"id": sms.ID,
		})
		return constants.ErrBackOffRetry
	}

	return err
}

func (c Consumer) sendOtp(_ context.Context, sms domain.SMS, message dto.Message, driver smsproviders.SmsSender) error {
	if message.Type == dto.TemplateTypeRow {
		if err := driver.SendSms(sms.Mobile, sms.Content); err != nil {
			return err
		}
		return nil
	}
	if len(message.Params) == 0 {
		return constants.ErrParamsAreNotProvided
	}
	code, err := strconv.Atoi(message.Params[0])
	if err != nil {
		c.logger.Error("failed to convert opt code to integer", log.J{
			"error": err.Error(),
		})
		return fmt.Errorf("failed to convert to integer: %w", err)
	}

	if err := driver.SendOTP(sms.Mobile, code); err != nil {
		return err
	}

	return nil
}

func (c Consumer) sendSms(_ context.Context, sms domain.SMS, message dto.Message, driver smsproviders.SmsSender) error {
	if message.Type == dto.TemplateTypeRow {
		if err := driver.SendSms(sms.Mobile, sms.Content); err != nil {
			return err
		}
		return nil
	}

	if err := driver.SendWithTemplate(sms.Mobile, message.Params...); err != nil {
		return err
	}

	return nil
}

func (c Consumer) isMessageExpired(ctx context.Context, sms domain.SMS) (bool, error) {
	if sms.ExpiresAt.After(time.Now()) {
		return false, nil
	}

	log.Warn("sms expiration reached before sending", log.J{
		"sms_id": sms.ID,
	})
	if err := c.smsRepository.UpdateStatus(ctx, sms.ID, domain.NotificationStatusExpired); err != nil {
		c.logger.Error("failed to update sms status", log.J{
			"error": err.Error(),
		})
		return true, err
	}

	return true, nil
}
