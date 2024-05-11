package transformer

import (
	"context"
	"notification/internal/domain"
)

type ProviderTemplateTransformer struct{}

type ResponseProviderTemplate struct {
	ID           uint   `json:"id" example:"1"`
	ProviderID   uint   `json:"provider_id" example:"2"`
	TemplateID   uint   `json:"template_id"`
	TemplateName string `json:"template_name" example:"122"`
}

func NewProviderTemplateTransformer() ProviderTemplateTransformer {
	return ProviderTemplateTransformer{}
}

func (p ProviderTemplateTransformer) Transform(_ context.Context, pt domain.ProviderTemplate) ResponseProviderTemplate {
	var respProviderTmpl ResponseProviderTemplate

	respProviderTmpl.ID = pt.ID
	respProviderTmpl.ProviderID = pt.ProviderID
	respProviderTmpl.TemplateID = pt.TemplateID
	respProviderTmpl.TemplateName = pt.Code

	return respProviderTmpl
}

func (p ProviderTemplateTransformer) TransformMany(ctx context.Context, pts []domain.ProviderTemplate) []ResponseProviderTemplate {
	var result []ResponseProviderTemplate
	for _, pr := range pts {
		transProvider := p.Transform(ctx, pr)

		result = append(result, transProvider)
	}

	return result
}
