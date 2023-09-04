package cidr_validation

import (
	"errors"
	"fmt"
	"net"
	"testing"
)

func TestCIDRContains(t *testing.T) {
	tests := []struct {
		inputs   []string // We could structure this so that that input 1 always is used as overarching addreses and see if all following addresses are contained in it.
		expected struct {
			returnString  string
			returnBoolean bool
			returnErr     error
		}
		expectedError error
	}{}
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
	inputs   []string
	expected struct {
		returnString  string
		returnBoolean bool
		returnErr     error
	}
	expectedError error
}) {
	for _, test := range tests {
		var testCaseString string
		var testCaseBool bool
		var err error

		testCaseString, testCaseBool, err = CIDRCompare(test.inputs[0], test.inputs[1:]...)
		fmt.Println(testCaseString)
		fmt.Println(testCaseBool)
		fmt.Println(err)
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
