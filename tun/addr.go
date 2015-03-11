package pelicantun

import (
	"fmt"
	"regexp"
	"strconv"
)

type addr struct {
	Ip     string
	Port   int
	IpPort string // Ip:Port
}

func NewAddr2(ip string, port int) addr {
	return addr{
		Ip:     ip,
		Port:   port,
		IpPort: fmt.Sprintf("%s:%d", ip, port),
	}
}

func (a *addr) SetIpPort() {
	a.IpPort = fmt.Sprintf("%s:%d", a.Ip, a.Port)
}

var ipColonPortRegex = regexp.MustCompile(`^([^:]*)\:(.*)$`)

func NewAddr1(ipport string) addr {
	var ip string
	var port int
	var err error

	match := ipColonPortRegex.FindStringSubmatch(ipport)
	//fmt.Printf("match = %#v\n  for ipport='%s'\n", match, ipport)
	if match == nil || len(match) != 3 {
		panic(fmt.Sprintf("bad scanf of ipport input '%s': could not find ip:port match", ipport))
	}
	ip = match[1]
	port, err = strconv.Atoi(match[2])
	panicOn(err)

	//fmt.Printf("ip = '%s', port = %d\n", ip, port)

	return addr{
		Ip:     ip,
		Port:   port,
		IpPort: ipport,
	}
}