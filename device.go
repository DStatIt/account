package account

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/satori/go.uuid"
)

const (
	devIDLength = 16
)

var src = rand.NewSource(time.Now().UnixNano())

type Device struct {
	GUID    string `json:"guid"`
	DevID   string `json:"device_id"`
	PhoneID string `json:"phone_id"`
}

func GenerateDevice() Device {
	return Device{
		GUID:    uuid.NewV4().String(),
		PhoneID: uuid.NewV4().String(),
		DevID:   generateDeviceID(),
	}
}

func generateDeviceID() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, devIDLength)
	for i := 0; i < devIDLength; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return fmt.Sprintf("android-%s", result)
}
