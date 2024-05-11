package domain

import "time"

type TemplatePriority string
type TemplateType string

const (
	TemplatePriorityHigh   TemplatePriority = "HIGH"
	TemplatePriorityMedium TemplatePriority = "MEDIUM"
	TemplatePriorityLow    TemplatePriority = "LOW"

	TemplateSMSType TemplateType = "SMS"
	TemplateOTPType TemplateType = "OTP"
)

type Template struct {
	ID                  uint
	Code                string
	Content             string
	Params              string
	UseProviderTemplate bool
	ActiveProviderID    uint
	Priority            TemplatePriority
	Type                TemplateType
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (t Template) IsDomain() bool {
	return true
}
