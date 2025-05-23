package cidr_validation

import (
	"errors"
	"testing"
)

func TestCheckCIDRsNotOverlap(t *testing.T) {
	tests := []struct {
		inputCIDRs    []string
		expected      bool
		expectedError error
	}{
		{
			[]string{"10.0.0.0/8"},
			false,
			ErrValidateInputCIDRs,
		},
		{
			[]string{"10.0.0.0/8", "10.0.0.0/16"},
			true,
			ErrValidateCIDRSCompare,
		},
		{
			[]string{"10.8.0.0/32", "10.8.0.0/32"},
			true,
			ErrValidateCIDRSCompare,
		},
		{
			[]string{"10.0.0.0/28", "10.0.0.0/28", "10.0.0.0/28"},
			true,
			ErrValidateCIDRSCompare,
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
		{
			[]string{" 10.8.0.16/28", " 10.8.0.0/28", "10.8.0.48/28 ", " 10.8.0.32/28 ", "10.8.0.64/28"},
			false,
			nil,
		},
	}
	runCheckCIDRsNotOverlap(t, tests)
}

func TestCheckCIDRsInNetworkRange(t *testing.T) {
	tests := []struct {
		inputNetwork  string
		inputCIDRs    []string
		expected      bool
		expectedError error
	}{
		// TODO: fix this test, maybe introduce a new error type
		{
			"bad data",
			[]string{"10.0.0.0/8"},
			false,
			ErrValidateNetworkRange,
		},
		{
			"10.0.0.0/8",
			[]string{"192.168.0.1/32"},
			false,
			ErrValidateCIDRSCompare,
		},
		{
			"10.0.0.0/8",
			[]string{"10.8.0.0/32", "192.168.0.1/32"},
			false,
			ErrValidateCIDRSCompare,
		},
		{
			"10.0.0.0/8",
			[]string{"10.0.0.0/28", "10.0.0.0/28", "192.168.0.1/32"},
			false,
			ErrValidateCIDRSCompare,
		},
		{
			"10.0.0.0/8",
			[]string{"10.0.0.0/24", "10.0.1.0/24"},
			true,
			nil,
		},
		{
			"10.0.0.0/8",
			[]string{"10.0.1.0/24", "10.0.2.0/24"},
			true,
			nil,
		},
		{
			"10.0.0.0/8",
			[]string{"10.8.0.0/28", "10.8.0.16/28"},
			true,
			nil,
		},
		{
			"10.0.0.0/8",
			[]string{"10.8.0.0/28", "10.8.0.16/28", "10.0.0.32/28"},
			true,
			nil,
		},
		{
			"10.0.0.0/8",
			[]string{"10.8.0.0/28", "10.8.0.16/28", "10.0.0.32/28", "10.0.0.48/28"},
			true,
			nil,
		},
		{
			"10.0.0.0/8",
			[]string{"10.8.0.16/28", "10.8.0.0/28", "10.0.0.48/28", "10.0.0.32/28"},
			true,
			nil,
		},
		{
			"10.0.0.0/8",
			[]string{"10.8.0.16/28", "10.8.0.0/28", "10.8.0.48/28", "10.8.0.32/28", "10.8.0.64/28"},
			true,
			nil,
		},
		{
			"10.0.0.0/8",
			[]string{"  10.8.0.16/28", "    10.8.0.0/28", "10.8.0.48/28   ", "  10.8.0.32/28   ", "  10.8.0.64/28  "},
			true,
			nil,
		},
		{
			"  10.0.0.0/8   ",
			[]string{"  10.8.0.16/28", "    10.8.0.0/28", "10.8.0.48/28   ", "  10.8.0.32/28   ", "  10.8.0.64/28  "},
			true,
			nil,
		},
	}
	runCheckCIDRsInNetworkRange(t, tests)
}

func runCheckCIDRsNotOverlap(t *testing.T, tests []struct {
	inputCIDRs    []string
	expected      bool
	expectedError error
}) {
	for _, test := range tests {
		var testCase bool
		var err error
		testCase, err = CheckCIDRsNotOverlap(test.inputCIDRs...)
		if testCase != test.expected {
			t.Logf("Test case: %+v\n", test)
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

func runCheckCIDRsInNetworkRange(t *testing.T, tests []struct {
	inputNetwork  string
	inputCIDRs    []string
	expected      bool
	expectedError error
}) {
	for _, test := range tests {
		var testCase bool
		var err error

		testCase, err = CheckCIDRsInNetworkRange(test.inputNetwork, test.inputCIDRs...)
		if testCase != test.expected {
			t.Logf("Test case: %+v\n", test)
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
