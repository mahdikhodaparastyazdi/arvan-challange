package domain

type ProviderTemplate struct {
	ID         uint
	ProviderID uint
	TemplateID uint
	Code       string
}

func (p ProviderTemplate) IsDomain() bool {
	return true
}
