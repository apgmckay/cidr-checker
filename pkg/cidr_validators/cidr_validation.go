package cidr_validation

import (
	"errors"
	"fmt"
	"net"
)

var ValidateCIDRSCompareErr = errors.New("ERROR: ip comparison failed")
var ValidateCIDRFailedErr = errors.New("ERROR: parsing cidr failed")
var ValidateInputCIDRsErr = errors.New("ERROR: must input 1 or more cidrs")

func CheckCIDRsNotOverlap(cidrAddrs ...string) (bool, error) {
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
				if ipnetA.Contains(ipnetB.IP) {
					result = true
					err = fmt.Errorf(
						"%w, IPs %s and %s are in the same range.\n",
						ValidateCIDRSCompareErr,
						cidrAddrs[i],
						cidrAddrs[j])
					return result, err
				}
			}
		}
	}
	return result, nil
}

func CheckCIDRsInNetworkRange(networkRange string, cidrAddrs ...string) (bool, error) {
	result := false

	_, ipnetNetworkRange, err := net.ParseCIDR(networkRange)
	if err != nil {
		return result, err
	}
	for i := range cidrAddrs {
		_, ipnet, err := net.ParseCIDR(cidrAddrs[i])
		if err != nil {
			return result, ValidateCIDRFailedErr
		}
		if ipnetNetworkRange.Contains(ipnet.IP) {
			result = true
			err = nil
		} else {
			result = false
			err = fmt.Errorf("%w, IPs %s and %s are not in the same range.\n", ValidateCIDRSCompareErr, ipnetNetworkRange, ipnet)
			return result, err
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
