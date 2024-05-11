package template

import (
	"context"
	"notification/internal/api/rest/requests"
	"notification/internal/domain"
)

func (s Service) UpdateTemplate(ctx context.Context, tmpl requests.UpdateTemplate) error {
	ut := domain.Template{
		ID:                  tmpl.ID,
		Code:                tmpl.Code,
		Content:             tmpl.Content,
		Params:              tmpl.Params,
		UseProviderTemplate: *tmpl.UseProviderTemplate,
		ActiveProviderID:    tmpl.ProviderID,
		Priority:            domain.TemplatePriority(tmpl.Priority),
		Type:                domain.TemplateType(tmpl.Type),
	}

	return s.TemplateRepository.Update(ctx, ut)
}
