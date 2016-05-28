package account

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/satori/go.uuid"
)

type AccountType int

const (
	Instagram   AccountType = iota
	devID       string      = "android-"
	devIDLength int         = 16 //DEV ID random string length
)

var src = rand.NewSource(time.Now().UnixNano())

type Account struct {
	Proxy Proxy

	ObjectID int `json:"_id"`

	Type AccountType `json:"account_type"`

	Username string `json:"username"`
	Password string `json:"password"`

	Email         string `json:"email"`
	EmailPassword string `json:"email_password"`

	PhoneNumber string `json:"phone_number"`

	GUID    string `json:"guid"`
	DevID   string `json:"device_id"`
	PhoneID string `json:"phone_id"`
}

func GenerateDeviceId() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, devIDLength)
	for i := 0; i < devIDLength; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return fmt.Sprintf("%s%s", devID, result)
}

func (a *Account) CreateDevice() {
	a.GUID = uuid.NewV4().String()
	a.PhoneID = uuid.NewV4().String()
	a.DevID = GenerateDeviceId()
}
