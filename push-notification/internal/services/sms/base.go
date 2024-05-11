package sms

import (
	"context"
	"time"
)

type smsRepository interface {
	CountSince(ctx context.Context, mobile string, since time.Time) (int, error)
}

type Service struct {
	smsRepository smsRepository
}

func New(
	sr smsRepository,
) Service {
	return Service{
		smsRepository: sr,
	}
}
