package cidr_validation

import (
	"errors"
	"net"
	"testing"
)

func TestCIDRContains(t *testing.T) {
	tests := []struct {
		inputs        []string
		expected      bool
		expectedError error
	}{
		{
			inputs:        []string{"10.0.0.0/8", "10.0.0.0/24"},
			expected:      true,
			expectedError: nil,
		},
		{

			inputs:        []string{"10.0.0.0/8", "10.0.0.0/24", "10.0.1.0/24"},
			expected:      true,
			expectedError: nil,
		},
		{

			inputs:        []string{"10.0.0.0/8", "10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24"},
			expected:      true,
			expectedError: nil,
		},
		{

			inputs:        []string{"10.0.0.0/8", "10.0.0.0/24", "192.168.0.1/16", "10.0.2.0/24"},
			expected:      false,
			expectedError: ValidateCIDRSCompareErr,
		},
		{

			inputs:        []string{"10.0.0.0/28", "10.2.0.0/24"},
			expected:      false,
			expectedError: ValidateCIDRSCompareErr,
		},
		{

			inputs:        []string{"10.0.0.0/28", "10.0.0.16/28"},
			expected:      false,
			expectedError: ValidateCIDRSCompareErr,
		},
		{

			inputs:        []string{"10.0.0.0/8", "10.0.0.0/28", "192.168.1.1/28", "10.0.1.0/28"},
			expected:      false,
			expectedError: ValidateCIDRSCompareErr,
		},
	}
	runCIDRContainsTests(t, tests)
}

func TestValidateCIDR(t *testing.T) {
	tests := []struct {
		inputs        []string
		expected      bool
		expectedError error
	}{
		{
			[]string{"10.0.0.0/8"},
			false,
			ValidateInputCIDRsErr,
		},
		{
			[]string{"10.0.0.0/8", "10.0.0.0/16"},
			true,
			ValidateCIDRSCompareErr,
		},
		{
			[]string{"10.8.0.0/32", "10.8.0.0/32"},
			true,
			ValidateCIDRSCompareErr,
		},
		{
			[]string{"10.0.0.0/24", "10.0.1.0/24"},
			false,
			nil,
		},
		{
			[]string{"10.0.1.0/24", "10.0.2.0/24"},
			false,
			nil,
		},
		{
			[]string{"10.8.0.0/28", "10.8.0.16/28"},
			false,
			nil,
		},
		{
			[]string{"10.8.0.0/28", "10.8.0.16/28", "10.0.0.32/28"},
			false,
			nil,
		},
		{
			[]string{"10.0.0.0/28", "10.0.0.0/28", "10.0.0.0/28"},
			true,
			ValidateCIDRSCompareErr,
		},
		{
			[]string{"10.8.0.0/28", "10.8.0.16/28", "10.0.0.32/28", "10.0.0.48/28"},
			false,
			nil,
		},
		{
			[]string{"10.8.0.16/28", "10.8.0.0/28", "10.0.0.48/28", "10.0.0.32/28"},
			false,
			nil,
		},
		{
			[]string{"10.8.0.16/28", "10.8.0.0/28", "10.8.0.48/28", "10.8.0.32/28", "10.8.0.64/28"},
			false,
			nil,
		},
	}
	runValidationTests(t, tests)
}

func runCIDRContainsTests(t *testing.T, tests []struct {
	inputs        []string
	expected      bool
	expectedError error
}) {
	for _, test := range tests {
		testCase, err := CIDRCompare(test.inputs[0], test.inputs[1:]...)
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

func runValidationTests(t *testing.T, tests []struct {
	inputs        []string
	expected      bool
	expectedError error
}) {
	for _, test := range tests {
		var testCase bool
		var err error
		testCase, err = ValidateCIDR(test.inputs...)
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
