package provider1

import (
	"notification/pkg/smsproviders"
)

type provider1 struct {
}

func NewProvider1() smsproviders.SmsSender {
	return provider1{}
}
