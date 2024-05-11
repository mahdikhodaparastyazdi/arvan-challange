package clientservice

import (
	"context"
	"notification/internal/domain"
)

type clientRepository interface {
	GetClientByApiKey(ctx context.Context, apiKey string) (domain.Client, error)
	UpdateBalance(ctx context.Context, amount int, clientId uint) error
}

type Service struct {
	clientRepository clientRepository
}

func New(cr clientRepository) Service {
	return Service{
		clientRepository: cr,
	}
}
