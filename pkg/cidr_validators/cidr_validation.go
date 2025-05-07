package cidr_validation

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

var ErrValidateNetworkRange = errors.New("ERROR: ip network range failed")
var ErrValidateCIDRSCompare = errors.New("ERROR: ip comparison failed")
var ErrValidateCIDRFailed = errors.New("ERROR: parsing cidr failed")
var ErrValidateInputCIDRs = errors.New("ERROR: must input 1 or more cidrs")

func CheckCIDRsNotOverlap(cidrAddrs ...string) (bool, error) {
	result, err := checkCIDRInputLength(cidrAddrs...)
	if err != nil {
		return result, err
	}

	var sanitizedCIDRAddrs []string
	sanitizedCIDRAddrs = sanitizeCIDRAddrs(cidrAddrs...)

	for i := range sanitizedCIDRAddrs {
		_, ipnetA, err := net.ParseCIDR(sanitizedCIDRAddrs[i])
		if err != nil {
			return result, ErrValidateCIDRFailed
		}
		for j := range cidrAddrs {
			if i != j {
				_, ipnetB, err := net.ParseCIDR(sanitizedCIDRAddrs[j])
				if err != nil {
					return result, ErrValidateCIDRFailed
				}
				if ipnetA.Contains(ipnetB.IP) {
					result = true
					err = fmt.Errorf(
						"%w, IPs %s and %s are in the same range.\n",
						ErrValidateCIDRSCompare,
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

	sanitizedNetworkRange := sanitizeNetworkAddr(networkRange)

	_, ipnetNetworkRange, err := net.ParseCIDR(sanitizedNetworkRange)
	if err != nil {
		return result, ErrValidateNetworkRange
	}

	var sanitizedCIDRAddrs []string
	sanitizedCIDRAddrs = sanitizeCIDRAddrs(cidrAddrs...)

	for i := range sanitizedCIDRAddrs {
		ip, _, err := net.ParseCIDR(sanitizedCIDRAddrs[i])
		if err != nil {
			return result, ErrValidateCIDRFailed
		}
		if ipnetNetworkRange.Contains(ip) {
			result = true
			err = nil
		} else {
			result = false
			err = fmt.Errorf("%w, IPs %s and %s are not in the same range.\n", ErrValidateCIDRSCompare, ipnetNetworkRange, ip)
			return result, err
		}
	}
	return result, err
}

func sanitizeNetworkAddr(networkRange string) string {
	return strings.TrimSpace(networkRange)
}

func sanitizeCIDRAddrs(cidrAddrs ...string) []string {
	var sanitizedCIDRAddrs []string
	for _, v := range cidrAddrs {
		sanitizedCIDRAddrs = append(sanitizedCIDRAddrs, strings.TrimSpace(v))
	}
	return sanitizedCIDRAddrs
}

func checkCIDRInputLength(cidrAddrs ...string) (bool, error) {
	result := false
	var err error
	cidrInputLength := len(cidrAddrs)
	if cidrInputLength <= 1 {
		result = false
		err = ErrValidateInputCIDRs
	}
	return result, err
}
