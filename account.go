package account

type AccountType int

const (
	Instagram AccountType = iota
	Facebook
)

type Account struct {
	ID int `sql:"id"`

	Type AccountType `sql:"account_type"`

	Username string `sql:"username"`
	Password string `sql:"password"`

	GUID    string `sql:"guid"`
	DevID   string `sql:"device_id"`
	PhoneID string `sql:"phone_id"`

	Proxy Proxy `sql:"proxy"`
}
