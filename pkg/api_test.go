package api

import (
	"fmt"
	"net"
	"testing"

	"github.com/3th1nk/cidr"
)

func TestCIDRCompare(t *testing.T) {
	_, ipnet, err := net.ParseCIDR("10.0.0.0/16")
	if err != nil {
		fmt.Println(err)
	}

	_, ipnet2, err := net.ParseCIDR("10.0.0.0/8")
	if err != nil {
		fmt.Println(err)
	}

	result := cidr.IPCompare(ipnet.IP, ipnet2.IP)
	switch result {
	case 0:
		println("same")
	case 1:
		println("less than")
	case -1:
		println("less than")
	default:
	}
}
