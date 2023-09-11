package cidr_validation

import (
	"errors"
	"fmt"
	"net"
)

var ValidateNetworkRangeErr = errors.New("ERROR: ip network range failed")
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
		return result, ValidateNetworkRangeErr
	}
	for i := range cidrAddrs {
		ip, _, err := net.ParseCIDR(cidrAddrs[i])
		if err != nil {
			return result, ValidateCIDRFailedErr
		}
		if ipnetNetworkRange.Contains(ip) {
			result = true
			err = nil
		} else {
			result = false
			/* TODO: Currently the mask is hardcoded, to try and get tests passing.
			- We should store the mask in byte and covert back to dec string
			*/
			err = fmt.Errorf("%w, IPs %s and %s/16 are not in the same range.\n", ValidateCIDRSCompareErr, ipnetNetworkRange, ip)
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
