package sms

import (
	"context"
	"time"
)

func (s Service) CountSince(ctx context.Context, mobile string, since time.Time) (int, error) {
	return s.smsRepository.CountSince(ctx, mobile, since)
}
