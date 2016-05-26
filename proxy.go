package account

import "fmt"

type Proxy struct {
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	InUse int    `json:"in_use"`
}

func (p Proxy) String() string {
	return fmt.Sprintf("http://%s:%d", p.IP, p.Port)
}
