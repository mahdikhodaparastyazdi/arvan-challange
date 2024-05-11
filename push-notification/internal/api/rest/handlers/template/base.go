package template

import (
	"context"
	"notification/internal/api/rest/requests"
	"notification/internal/api/rest/transformer"
	"notification/internal/domain"

	response_formatter "notification/pkg/response_formatter"
)

type templateServiceInterface interface {
	SendTemplate(ctx context.Context, msg requests.SendWithTemplateRequest, clientId uint) error
	GetAllTemplates(ctx context.Context) ([]domain.Template, error)
	CreteTemplate(ctx context.Context, tmpl requests.CreateTemplate) (domain.Template, error)
	UpdateTemplate(ctx context.Context, tmpl requests.UpdateTemplate) error
}

type Handler struct {
	templateService     templateServiceInterface
	responseFormatter   response_formatter.ResponseFormatter
	templateTransformer transformer.TemplateTransformer
}

func New(templateService templateServiceInterface,
	responseFormatter response_formatter.ResponseFormatter,
	templateTrans transformer.TemplateTransformer,
) Handler {
	return Handler{
		templateService:     templateService,
		responseFormatter:   responseFormatter,
		templateTransformer: templateTrans,
	}
}
