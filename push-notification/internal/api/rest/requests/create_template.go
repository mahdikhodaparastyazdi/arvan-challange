package requests

type CreateTemplate struct {
	Code                string `json:"code" validate:"required"`
	Content             string `json:"content" validate:"required"`
	Params              string `json:"params" validate:"required"`
	ProviderID          uint   `json:"provider_id" validate:"required,number"`
	UseProviderTemplate *bool  `json:"use_provider_template" validate:"required,boolean"`
	Priority            string `json:"priority" validate:"required"`
	Type                string `json:"type" validate:"required"`
}

func (c *CreateTemplate) PrepareForValidation() {}
