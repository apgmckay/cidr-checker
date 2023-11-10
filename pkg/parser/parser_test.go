package parser

import (
	cidr_validatiors "cidr-checker/pkg/cidr_validators"
	"errors"
	"testing"
)

func TestParserAndRun(t *testing.T) {
	helpOutput := PrintHelpOutput()
	outputSuccessContains, _ := successOutput("contains")
	outputSuccessNoOverlap, _ := successOutput("no-overlap")

	tests := []struct {
		inputs        []string
		expected      string
		expectedError error
	}{
		{
			[]string{"--help"},
			helpOutput,
			nil,
		},
		{
			[]string{"   --help    "},
			helpOutput,
			nil,
		},
		{
			[]string{"10.0.0.0/24", "10.0.1.0/24"},
			outputSuccessNoOverlap,
			nil,
		},
		{
			[]string{"10.0.0.0/24", "10.0.0.0/24"},
			cidr_validatiors.ValidateCIDRSCompareErr.Error(),
			cidr_validatiors.ValidateCIDRSCompareErr,
		},
		{
			[]string{"--network", "10.0.0.0/8", "10.0.0.0/24"},
			outputSuccessContains,
			nil,
		},
		{
			[]string{"--network", "10.0.0.0/8", "10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"},
			outputSuccessContains,
			nil,
		},
		{
			[]string{"--network", "10.0.0.0/8", "10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24", "192.168.1.0/16"},
			"",
			cidr_validatiors.ValidateCIDRSCompareErr,
		},
		{
			[]string{"--network", " 10.0.0.0/8", "  10.0.0.0/24", "10.0.1.0/24  ", "  10.0.2.0/24  "},
			outputSuccessContains,
			nil,
		},
		{
			[]string{"--network", " 10.0.0.0/8", "  10.0.0.0/24", "10.0.1.0/24  ", "  10.0.2.0/24  "},
			outputSuccessContains,
			nil,
		},
		{
			[]string{"  --network", " 10.0.0.0/8", "  10.0.0.0/24", "10.0.1.0/24  ", "  10.0.2.0/24  "},
			outputSuccessContains,
			nil,
		},
		{
			[]string{"--network"},
			helpOutput,
			ParseErr,
		},
	}
	runTests(t, tests)
}

func runTests(t *testing.T, tests []struct {
	inputs        []string
	expected      string
	expectedError error
}) {
	for _, test := range tests {
		var testCase string
		var err error
		testCase, err = ParseAndRun(test.inputs...)
		if testCase != test.expected {
			t.Logf("testCase got: %s \n\texpected: %s.\n", testCase, test.expected)
			t.Logf("Failing!")
			t.Fail()
		}
		if errors.Is(err, test.expectedError) {
			t.Logf("Errors are equal.\n")
		} else {
			t.Logf("Errors are not equal.\n\texpected: %s\n\tgot     : %s", test.expectedError, err)
			t.Fail()
			t.Logf("Failing!")
		}
	}

}
