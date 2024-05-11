package provider2

import (
	"notification/pkg/smsproviders"
)

type provider2 struct {
}

func NewProvider2() smsproviders.SmsSender {
	return provider2{}
}
