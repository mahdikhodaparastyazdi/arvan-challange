package domain

type NotificationStatus string

const (
	NotificationStatusPending NotificationStatus = "PENDING"
	NotificationStatusSent    NotificationStatus = "SENT"
	NotificationStatusRetry   NotificationStatus = "RETRY"
	NotificationStatusFailed  NotificationStatus = "FAILED"
	NotificationStatusExpired NotificationStatus = "EXPIRED"
)

func (n NotificationStatus) String() string {
	return string(n)
}
