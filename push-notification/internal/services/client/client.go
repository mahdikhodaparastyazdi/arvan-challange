package clientservice

import (
	"context"
	"notification/internal/domain"
)

func (s Service) GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error) {
	client, err := s.clientRepository.GetClientByApiKey(ctx, apiKey)
	if err != nil {
		return client, err
	}

	return client, nil
}
func (s Service) UpdateBalance(ctx context.Context, amount int, clientId uint) error {
	return s.clientRepository.UpdateBalance(ctx, amount, clientId)
}
