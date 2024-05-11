package template

import (
	"context"
	"notification/internal/domain"
)

func (s Service) GetAllTemplates(ctx context.Context) ([]domain.Template, error) {
	return s.TemplateRepository.GetAll(ctx)
}
