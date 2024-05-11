package transformer

import (
	"context"
	"notification/internal/domain"
)

type ProviderTransformer struct{}

type ResponseProvider struct {
	ID     uint   `json:"id" example:"1"`
	Name   string `json:"name" example:"discord"`
	Active bool   `json:"active" example:"true"`
}

func NewProviderTransformer() ProviderTransformer {
	return ProviderTransformer{}
}

func (p ProviderTransformer) Transform(_ context.Context, provider domain.Provider) ResponseProvider {
	var respProvider ResponseProvider

	respProvider.ID = provider.ID
	respProvider.Active = provider.Active
	respProvider.Name = provider.Name
	return respProvider
}

func (p ProviderTransformer) TransformMany(ctx context.Context, providers []domain.Provider) []ResponseProvider {
	var result []ResponseProvider
	for _, pr := range providers {
		transProvider := p.Transform(ctx, pr)

		result = append(result, transProvider)
	}

	return result
}
