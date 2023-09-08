package cidr_validation

import (
	"errors"
	"fmt"
	"net"
)

var ValidateCIDRSCompareErr = errors.New("ERROR: ip comparison failed")
var ValidateCIDRFailedErr = errors.New("ERROR: parsing cidr failed")

func ValidateCIDR(errorOnMatch bool, cidrAddrs ...string) (bool, error) {
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
					if errorOnMatch {
						err = fmt.Errorf(
							"%w, IPs %s and %s are in the same range.\n",
							ValidateCIDRSCompareErr,
							cidrAddrs[i],
							cidrAddrs[j])
					} else {
						err = nil
					}
					return result, err
				}
			}
		}
	}
	return result, nil
}

func checkCIDRInputLength(cidrAddrs ...string) (bool, error) {
	result := false
	var err error
	if len(cidrAddrs) <= 1 {
		result = false
		err = ValidateCIDRFailedErr
	}
	return result, err
}
