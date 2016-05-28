package account

type AccountType int

const (
	Instagram AccountType = iota
)

type Account struct {
	Proxy Proxy

	ObjectID int `json:"_id"`

	Type AccountType `json:"account_type"`

	Username string `json:"username"`
	Password string `json:"password"`

	Email         string `json:"email"`
	EmailPassword string `json:"email_password"`

	PhoneNumber string `json:"phone_number"`

	Device Device `json:"device"`

	Cookies []Cookie `json:"cookies"`
}
