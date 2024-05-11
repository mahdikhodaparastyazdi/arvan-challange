package requests

type CreateProviderTemplate struct {
	ProviderID   uint   `json:"provider_id" validate:"required"`
	TemplateID   uint   `json:"template_id" validate:"required"`
	TemplateName string `json:"template_name" validate:"required"`
}

func (c *CreateProviderTemplate) PrepareForValidation() {}
