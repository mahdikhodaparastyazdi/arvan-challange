package domain

type Client struct {
	ID       uint
	Balance  int
	SMSRate  string
	Username string
	Token    string
	IsActive bool
}

func (c Client) IsDomain() bool {
	return true
}
