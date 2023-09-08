package cidr_validation

import (
	"errors"
	"net"
	"testing"
)

func TestValidateCIDR(t *testing.T) {
	tests := []struct {
		inputBool     bool
		inputCIDRs    []string
		expected      bool
		expectedError error
	}{
		{
			true,
			[]string{""},
			false,
			ValidateCIDRFailedErr,
		},
		{
			false,
			[]string{""},
			false,
			ValidateCIDRFailedErr,
		},
		{
			true,
			[]string{"10.0.0.0/8"},
			false,
			ValidateCIDRFailedErr,
		},
		{
			true,
			[]string{"bad data"},
			false,
			ValidateCIDRFailedErr,
		},
		{
			true,
			[]string{"10.0.0.0/8", "10.0.0.0/16"},
			true,
			ValidateCIDRSCompareErr,
		},
		{
			true,
			[]string{"10.8.0.0/32", "10.8.0.0/32"},
			true,
			ValidateCIDRSCompareErr,
		},
		{
			true,
			[]string{"10.0.0.0/28", "10.0.0.0/28", "10.0.0.0/28"},
			true,
			ValidateCIDRSCompareErr,
		},
		{
			true,
			[]string{"10.0.0.0/24", "10.0.1.0/24"},
			false,
			nil,
		},
		{
			true,
			[]string{"10.0.1.0/24", "10.0.2.0/24"},
			false,
			nil,
		},
		{
			true,
			[]string{"10.8.0.0/28", "10.8.0.16/28"},
			false,
			nil,
		},
		{
			true,
			[]string{"10.8.0.0/28", "10.8.0.16/28", "10.0.0.32/28"},
			false,
			nil,
		},
		{
			true,
			[]string{"10.8.0.0/28", "10.8.0.16/28", "10.0.0.32/28", "10.0.0.48/28"},
			false,
			nil,
		},
		{
			true,
			[]string{"10.8.0.16/28", "10.8.0.0/28", "10.0.0.48/28", "10.0.0.32/28"},
			false,
			nil,
		},
		{
			true,
			[]string{"10.8.0.16/28", "10.8.0.0/28", "10.8.0.48/28", "10.8.0.32/28", "10.8.0.64/28"},
			false,
			nil,
		},
		{
			false,
			[]string{"10.0.0.0/8", "10.8.0.0/28", "10.8.0.48/28", "10.8.0.32/28", "10.8.0.64/28"},
			true,
			nil,
		},
	}
	runValidationTests(t, tests)
}

func runValidationTests(t *testing.T, tests []struct {
	inputBool     bool
	inputCIDRs    []string
	expected      bool
	expectedError error
}) {
	for _, test := range tests {
		var testCase bool
		var err error
		testCase, err = ValidateCIDR(test.inputBool, test.inputCIDRs...)
		if testCase != test.expected {
			t.Logf("%t and %t are not equal.\n", testCase, test.expected)
			t.Fail()
		}
		if errors.Is(err, test.expectedError) {
			t.Logf("Errors are equal.\n")
		} else {
			t.Logf("Errors are not equal.\n\texpected: %s\n\tgot: %s", test.expectedError, err)
			t.Fail()
		}
	}

}

func testParseIP(input string) string {
	_, ipNet, _ := net.ParseCIDR(input)
	return ipNet.String()
}
