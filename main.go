package main

import (
	cidr_validator "cidr-checker/pkg/cidr_validators"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		result, err := cidr_validator.ValidateCIDR(os.Args[1:]...)
		if result {
			fmt.Printf("%s\n", err)
		} else {
			fmt.Printf("All good no overlapping CIDRs.\n")
		}
	} else {
		// TODO: we also want to read from Stdin to allow piping to the program
	}
}
