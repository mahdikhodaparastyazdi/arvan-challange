package requests

type SendWithTemplateRequest struct {
	Mobile       string            `json:"mobile" validate:"required"`
	TemplateCode string            `json:"template_code" validate:"required"`
	Params       map[string]string `json:"params"`
	ExpiresAt    string            `json:"expires_at" validate:"omitempty,time"`
}

func (s *SendWithTemplateRequest) PrepareForValidation() {}
