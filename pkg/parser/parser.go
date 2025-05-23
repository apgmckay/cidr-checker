package parser

import (
	cidr_validators "cidr-checker/pkg/cidr_validators"
	"errors"
	"fmt"
	"strings"
)

var HelpStdOutputLineA = fmt.Sprintln("\ncidr-check recieved No input.")
var HelpStdOutputLineB = fmt.Sprintln("Entered from:")
var HelpStdOutputLineC = fmt.Sprintln("")
var HelpStdOutputLineD = fmt.Sprintln("os.Args like `cidr-checker 10.0.0.0/19 10.0.1.0/19 10.0.2.0/19`.")
var HelpStdOutputLineE = fmt.Sprintln("--network\tcan be used to compare given addresses to the value of network.\n\t\tfor example: `cidr-checker 10.0.0.0/24 10.0.0.1/24 --network 10.0.0.0/8`")

var HelpStdOutput = fmt.Sprintf("%s%s%s%s%s\n",
	HelpStdOutputLineA,
	HelpStdOutputLineB,
	HelpStdOutputLineC,
	HelpStdOutputLineD,
	HelpStdOutputLineE)

var ErrParse = errors.New("Parseing Error.\n")

type InputFlags struct {
	NetworkAddr string
	HelpSet     bool
}

func ParseAndRun(input ...string) (string, error) {
	var output string

	flags := NewInputFlags()

	parsedInput, err := flags.Parse(input...)
	if err != nil {
		return PrintHelpOutput(), err
	}

	if flags.HelpSet {
		output = PrintHelpOutput()
		return output, err
	}
	if len(flags.NetworkAddr) >= 1 {
		result, err := cidr_validators.CheckCIDRsInNetworkRange(flags.NetworkAddr, parsedInput...)
		if !result {
			return "", err
		} else {
			output, err = successOutput("contains")
			return output, err
		}
	}
	result, err := cidr_validators.CheckCIDRsNotOverlap(parsedInput...)
	if !result {
		output, err = successOutput("no-overlap")
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

			sanitizedInput := strings.TrimSpace(input[i])

			switch sanitizedInput {
			case "--help":
				f.HelpSet = true
				return []string{}, nil
			case "--network":
				if len(input) < 2 {
					return output, ErrParse
				}
				f.SetNetworkAddr(input[i+1])
				i++
			default:
				output = append(output, sanitizedInput)
			}
		}
	} else {
		f.HelpSet = true
	}
	return output, err
}

func PrintHelpOutput() string {
	return HelpStdOutput
}

func successOutput(input string) (string, error) {
	var output string
	var err error
	switch input {
	case "contains":
		output = "All good all ips in network range.\n"
	case "no-overlap":
		output = "All good no overlapping CIDRs.\n"
	default:
		// error here
	}
	return output, err
}
