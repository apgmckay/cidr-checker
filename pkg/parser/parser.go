package parser

import (
	cidr_validators "cidr-checker/pkg/cidr_validators"
	"errors"
	"fmt"
)

var ParseErr = errors.New("Parseing Error.\n")

type InputFlags struct {
	NetworkAddr string
	HelpSet     bool
}

func ParseAndRun(input ...string) (string, error) {
	var output string
	var err error

	flags := NewInputFlags()

	parsedInput, err := flags.Parse(input...)
	if err != nil {
		return helpOutput(), err
	}

	if flags.HelpSet {
		output = helpOutput()
		return output, err
	}
	if len(flags.NetworkAddr) >= 1 {
		result, err := cidr_validators.CheckCIDRsInNetworkRange(flags.NetworkAddr, parsedInput...)
		if result == false {
			return "", err
		} else {
			return fmt.Sprintf("All good all ips in network range.\n"), err
		}
	}
	result, err := cidr_validators.CheckCIDRsNotOverlap(parsedInput...)
	if result == false {
		output = fmt.Sprintf("All good no overlapping CIDRs.\n")
		err = nil
	} else {
		output = fmt.Sprintf("%s", errors.Unwrap(err))
	}
	return output, err
}

func NewInputFlags() *InputFlags {
	return &InputFlags{}
}

func (f *InputFlags) SetNetworkAddr(input string) {
	f.NetworkAddr = input
}

func (f *InputFlags) SetHelp() {
	f.HelpSet = true
}

func (f *InputFlags) Parse(input ...string) ([]string, error) {
	var output []string
	var err error
	if len(input) >= 1 {
		for i := range input {
			switch input[i] {
			case "--help":
				f.HelpSet = true
				return []string{}, nil
			case "--network":
				f.SetNetworkAddr(input[i+1])
				i++
			default:
				output = append(output, input[i])
			}
		}
	} else {
		f.HelpSet = true
	}
	return output, err
}

func helpOutput() string {
	lineA := fmt.Sprintln("\ncidr-check recieved No input.")
	lineB := fmt.Sprintln("Entered from:")
	lineC := fmt.Sprintln("os.Args like `cidr-checker 10.0.0.0/19 10.0.1.0/19 10.0.2.0/19`.")
	lineD := fmt.Sprintln("--network\tcan be used to compare given addresses to the value of network.\n\t\tfor example: `cidr-checker 10.0.0.0/24 10.0.0.1/24 --network 10.0.0.0/8`")
	return fmt.Sprintf("%s%s%s%s", lineA, lineB, lineC, lineD)

}
