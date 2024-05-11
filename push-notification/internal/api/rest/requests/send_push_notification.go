package requests

type PushNotificationPayload struct {
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	ImageUrl string `json:"image_url"`
	DeepLink string `json:"deep_link"`
}
type PushNotificationRequest struct {
	UserID    uint                    `json:"user_id" validate:"required"`
	Payload   PushNotificationPayload `json:"payload" validate:"required"`
	ExpiresAt string                  `json:"expires_at" validate:"omitempty,time"`
}

func (p *PushNotificationRequest) PrepareForValidation() {}
