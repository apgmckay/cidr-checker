package parser

import (
	"errors"
	"fmt"
	"testing"
)

func TestParserAndRun(t *testing.T) {
	lineA := fmt.Sprintln("\ncidr-check recieved No input.")
	lineB := fmt.Sprintln("Entered from:")
	lineC := fmt.Sprintln("os.Args like `cidr-checker 10.0.0.0/19 10.0.1.0/19 10.0.2.0/19`.")
	output := fmt.Sprintf("%s%s%s", lineA, lineB, lineC)

	outputSuccess := fmt.Sprintf("All good no overlapping CIDRs.\n")

	tests := []struct {
		inputs        []string
		expected      string
		expectedError error
	}{
		{
			[]string{},
			output,
			parseErr,
		},
		{
			[]string{"10.0.0.0/24", " 10.0.1.0/24"},
			outputSuccess,
			nil,
		},
		// TODO: this is currently comment out figure this test out
		//		{
		//			[]string{"10.0.0.0/24", " 10.0.0.0/24"},
		//			cidr_validator.ValidateCIDRSCompareErr.Error(),
		//			cidr_validator.ValidateCIDRSCompareErr,
		//		},
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
			t.Logf("got: %s \nexpected: %s.\n", testCase, test.expected)
			t.Logf("Failing!")
			t.Fail()
		}
		if errors.Is(err, test.expectedError) {
			t.Logf("Errors are equal.\n")
		} else {
			t.Logf("Errors are not equal.\n\texpected: %s\n\tgot: %s", test.expectedError, err)
			t.Fail()
			t.Logf("Failing!")
		}
	}

}
