package dto

type TemplateType string

const (
	TemplateTypeRow      TemplateType = "ROW"
	TemplateTypeTemplate TemplateType = "TEMPLATE"
)

type Message struct {
	Id                   uint         `json:"id"`
	ProviderID           uint         `json:"provider_id"`
	ProviderName         string       `json:"provider_name"`
	ProviderTemplateCode string       `json:"provider_template_code"`
	SmsType              string       `json:"sms_type"`
	Type                 TemplateType `json:"type"`
	Params               []string     `json:"params"`
}
