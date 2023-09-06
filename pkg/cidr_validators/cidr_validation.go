package cidr_validation

import (
	"errors"
	"fmt"
	"net"

	"github.com/3th1nk/cidr"
)

var ValidateCIDRSCompareErr = errors.New("ERROR: ip comparison failed")
var ValidateCIDRFailedErr = errors.New("ERROR: parsing cidr failed")
var ValidateInputCIDRsErr = errors.New("ERROR: must input 1 or more cidrs")

func ValidateCIDR(cidrAddrs ...string) (bool, error) {
	result, err := checkCIDRInputLength(cidrAddrs...)
	if err != nil {
		return result, err
	}
	for i := range cidrAddrs {
		_, ipnetA, err := net.ParseCIDR(cidrAddrs[i])
		if err != nil {
			return result, ValidateCIDRFailedErr
		}
		for j := range cidrAddrs {
			if i != j {
				_, ipnetB, err := net.ParseCIDR(cidrAddrs[j])
				if err != nil {
					return result, ValidateCIDRFailedErr
				}
				compareResult := cidr.IPCompare(ipnetA.IP, ipnetB.IP)
				switch compareResult {
				case 0:
					result = true
					return result, fmt.Errorf("%w, IPs %s and %s are in the same range.\n", ValidateCIDRSCompareErr, cidrAddrs[i], cidrAddrs[j])
				}
			}
		}
	}
	return result, nil
}

func CIDRCompare(candidate string, cidrAddrs ...string) (bool, error) {
	var result bool
	_, ipnetA, err := net.ParseCIDR(candidate)
	if err != nil {
		return false, err
	}
	for _, v := range cidrAddrs {
		_, ipnetB, err := net.ParseCIDR(v)
		if err != nil {
			return false, err
		}

		if ipnetA.Contains(ipnetB.IP) {
			result = true
			err = nil
		} else {
			return false, ValidateCIDRSCompareErr
		}
	}
	return result, err
}

func checkCIDRInputLength(cidrAddrs ...string) (bool, error) {
	result := false
	var err error
	cidrInputLength := len(cidrAddrs)
	if cidrInputLength <= 1 {
		result = false
		err = ValidateInputCIDRsErr
	}
	return result, err
}
