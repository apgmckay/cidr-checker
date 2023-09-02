package parser

import (
	cidr_validator "cidr-checker/pkg/cidr_validators"
	"errors"
	"fmt"
)

var parseErr = errors.New("Parseing Error.\n")

func ParseAndRun(input ...string) (string, error) {
	var output string
	var err error
	if len(input) > 1 {
		result, err := cidr_validator.ValidateCIDR(input...)
		if result {
			output = fmt.Sprintf("%s", err)
		} else {
			output = fmt.Sprintf("All good no overlapping CIDRs.\n")
		}
	} else {
		lineA := fmt.Sprintln("\ncidr-check recieved No input.")
		lineB := fmt.Sprintln("Entered from:")
		lineC := fmt.Sprintln("os.Args like `cidr-checker 10.0.0.0/19 10.0.1.0/19 10.0.2.0/19`.")
		output = fmt.Sprintf("%s%s%s", lineA, lineB, lineC)
		err = parseErr
	}
	return output, err
}
