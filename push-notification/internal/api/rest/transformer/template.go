package transformer

import (
	"context"
	"notification/internal/domain"
)

type TemplateTransformer struct{}

type ResponseTemplate struct {
	ID                  uint   `json:"id" example:"1"`
	Code                string `json:"code" example:"42"`
	Content             string `json:"content"`
	Params              string `json:"params" example:"username,age"`
	UseProviderTemplate bool   `json:"use_provider_template" example:"true"`
	ActiveProviderID    uint   `json:"active_provider_id" example:"4"`
	Priority            string `json:"priority" example:"HIGH"`
	Type                string `json:"type" example:"OTP"`
}

func NewTemplateTransformer() TemplateTransformer {
	return TemplateTransformer{}
}

func (t TemplateTransformer) Transform(_ context.Context, tmpl domain.Template) ResponseTemplate {
	var resp ResponseTemplate

	resp.ID = tmpl.ID
	resp.UseProviderTemplate = tmpl.UseProviderTemplate
	resp.ActiveProviderID = tmpl.ActiveProviderID
	resp.Content = tmpl.Content
	resp.Code = tmpl.Code
	resp.Priority = string(tmpl.Priority)
	resp.Type = string(tmpl.Type)
	resp.Params = tmpl.Params

	return resp
}

func (t TemplateTransformer) TransformMany(ctx context.Context, templates []domain.Template) []ResponseTemplate {
	var result []ResponseTemplate
	for _, tm := range templates {
		transProvider := t.Transform(ctx, tm)

		result = append(result, transProvider)
	}

	return result
}
