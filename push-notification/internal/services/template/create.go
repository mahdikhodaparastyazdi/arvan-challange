package template

import (
	"context"
	"notification/internal/api/rest/requests"
	"notification/internal/domain"
)

func (s Service) CreteTemplate(ctx context.Context, tmpl requests.CreateTemplate) (domain.Template, error) {
	newTemplate := domain.Template{
		Code:                tmpl.Code,
		Content:             tmpl.Content,
		Params:              tmpl.Params,
		UseProviderTemplate: *tmpl.UseProviderTemplate,
		ActiveProviderID:    tmpl.ProviderID,
		Priority:            domain.TemplatePriority(tmpl.Priority),
		Type:                domain.TemplateType(tmpl.Type),
	}

	d, err := s.TemplateRepository.Create(ctx, newTemplate)
	if err != nil {
		return domain.Template{}, err
	}

	newTemplate.ID = d.ID

	return newTemplate, nil
}
