package account

import "fmt"

type Proxy struct {
	IP   string `sql:"ip"`
	Port int    `sql:"port"`
}

func (p Proxy) String() string {
	return fmt.Sprintf("http://%s:%d", p.IP, p.Port)
}
