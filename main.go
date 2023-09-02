package main

import (
	parser "cidr-checker/pkg/parser"
	"log"
	"os"
)

func main() {
	output, err := parser.ParseAndRun(os.Args[1:]...)
	if err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
	log.Printf("%s", output)
}
