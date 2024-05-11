package smsproviders

type SmsSender interface {
	SendOTP(receptor string, otpCode int) error
	SendSms(receptor, text string) error
	SendWithTemplate(receptor string, params ...string) error
}
