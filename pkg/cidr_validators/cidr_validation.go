package cidr_validation

import (
	"errors"
	"net"

	"github.com/3th1nk/cidr"
)

var validateCIDRSCompareErr = errors.New("ERROR: ip comparison failed")
var validateCIDRFailedErr = errors.New("ERROR: parsing cidr failed")
var validateInputCIDRsErr = errors.New("ERROR: must input 1 or more cidrs")

func ValidateCIDR(cidrAddrs ...string) (bool, error) {
	result, err := checkCIDRInputLength(cidrAddrs...)
	if err != nil {
		return result, err
	}
	for i := range cidrAddrs {
		_, ipnetA, err := net.ParseCIDR(cidrAddrs[i])
		if err != nil {
			return result, validateCIDRFailedErr
		}
		for j := range cidrAddrs {
			if i != j {
				_, ipnetB, err := net.ParseCIDR(cidrAddrs[j])
				if err != nil {
					return result, validateCIDRFailedErr
				}
				compareResult := cidr.IPCompare(ipnetA.IP, ipnetB.IP)
				switch compareResult {
				case 0:
					result = true
					return result, validateCIDRSCompareErr

				}
			}
		}
	}
	return result, nil
}

func checkCIDRInputLength(cidrAddrs ...string) (bool, error) {
	result := false
	var err error
	cidrInputLength := len(cidrAddrs)
	if cidrInputLength <= 1 {
		result = false
		err = validateInputCIDRsErr
	}
	return result, err
}
