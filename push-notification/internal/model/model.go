package model

import "notification/internal/domain"

type DbModel interface {
	TableName() string
	ToDomain() domain.Domain
}
