package template

import (
	"context"
	"fmt"
	"notification/internal/api/rest/requests"
	"notification/internal/constants"
	"notification/internal/domain"
	"notification/internal/dto"
	"strings"
	"time"

	log "notification/pkg/logger"
)

func (s Service) SendTemplate(ctx context.Context, req requests.SendWithTemplateRequest, clinetId uint) error {

	err := s.ClientRepository.UpdateBalance(ctx, constants.SmsPrice, clinetId)
	if err != nil {
		return err
	}
	tmpl, err := s.TemplateRepository.FindByCode(ctx, req.TemplateCode)
	if err != nil {
		return err
	}

	p, err := s.providerRepository.ByID(ctx, tmpl.ActiveProviderID)
	if err != nil {
		return err
	}

	if !p.Active {
		s.logger.Info("provider is not active", log.J{
			"name": p.Name,
			"id":   p.ID,
		})
		return constants.ErrProviderIsNotActive
	}

	if !tmpl.UseProviderTemplate {
		return s.SendWithRowTemplate(ctx, req, tmpl, &p)
	}

	return s.SendWithProviderTemplate(ctx, req, tmpl, &p)
}

func (s Service) SendWithRowTemplate(ctx context.Context, req requests.SendWithTemplateRequest, tmpl domain.Template, provider *domain.Provider) error {
	content := s.createContent(tmpl.Content, req.Params)

	params := make([]string, 0)
	if len(tmpl.Params) != 0 {
		templateParams := strings.Split(tmpl.Params, ",")
		for _, k := range templateParams {
			if _, ok := req.Params[k]; !ok {
				return constants.ErrTemplateParamsIsRequired
			}
			params = append(params, req.Params[k])
		}
	}

	var exp time.Time
	if req.ExpiresAt != "" {
		exp, _ = time.Parse(time.RFC3339, req.ExpiresAt)
	} else {
		exp = time.Now()
	}

	if tmpl.Type == constants.SMS_TYPE_SMS {
		exp = exp.Add(s.messagesTTL.SMS)
	} else {
		exp = exp.Add(s.messagesTTL.OTP)
	}

	sms := domain.SMS{
		ProviderId: tmpl.ActiveProviderID,
		TemplateID: tmpl.ID,
		Mobile:     req.Mobile,
		Content:    content,
		Status:     domain.NotificationStatusPending,
		ExpiresAt:  exp,
	}

	sms, err := s.SmsRepository.Create(ctx, sms)
	if err != nil {
		return err
	}

	priority := s.getTemplatePriority(tmpl.Priority)
	if err != nil {
		return err
	}

	t := dto.Message{
		Id:           sms.ID,
		ProviderID:   provider.ID,
		ProviderName: provider.Name,
		SmsType:      string(tmpl.Type),
		Type:         dto.TemplateTypeRow,
		Params:       params,
	}

	return s.queue.Enqueue(t, priority)
}

func (s Service) SendWithProviderTemplate(ctx context.Context, req requests.SendWithTemplateRequest, tmpl domain.Template, provider *domain.Provider) error {
	// TODO: implement send with provider template
	return nil
}

func (s Service) getTemplatePriority(p domain.TemplatePriority) string {
	switch p {
	case domain.TemplatePriorityHigh:
		return constants.QUEUE_PRIORITY_HIGH
	case domain.TemplatePriorityLow:
		return constants.QUEUE_PRIORITY_LOW
	default:
		return constants.QUEUE_PRIORITY_LOW
	}
}

func (s Service) createContent(content string, params map[string]string) string {
	for key, value := range params {
		placeholder := fmt.Sprintf("{%s}", key)
		content = strings.ReplaceAll(content, placeholder, value)
	}

	return content
}
