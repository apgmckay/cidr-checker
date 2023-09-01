package cidr_validation

import (
	"errors"
	"net"

	"github.com/3th1nk/cidr"
)

var validateCIDRSCompareErr = errors.New("ERROR: ip comparison failed")
var validateCIDRFailedErr = errors.New("ERROR: parsing cidr failed")
var validateInputCIDRsErr = errors.New("ERROR: must input 1 or more cidrs")

type parsedCIDRs struct {
	ipnets []*net.IPNet
	errs   []error
}

func validateCIDR(cidrAddrs ...string) (bool, error) {
	cidrInputLength := len(cidrAddrs)
	if cidrInputLength <= 1 {
		return false, validateInputCIDRsErr
	}
	var pCIDRs parsedCIDRs
	for i := 0; i < cidrInputLength; i++ {
		_, ipnet, err := net.ParseCIDR(cidrAddrs[i])
		if err != nil {

			return false, validateCIDRFailedErr

		}
		pCIDRs.ipnets = append(pCIDRs.ipnets, ipnet)
		pCIDRs.errs = append(pCIDRs.errs, err)
	}

	for j := 1; j < len(pCIDRs.ipnets); j++ {
		result := cidr.IPCompare(pCIDRs.ipnets[j-1].IP, pCIDRs.ipnets[j].IP)
		switch result {
		case 0:

			return true, validateCIDRSCompareErr

		}
	}

	return false, nil
}
