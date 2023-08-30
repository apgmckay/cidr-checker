package main

import (
	"io"
	"os"
)

func main() {
	//	bytes, _ := io.ReadAll(os.Stdin)
	if len(os.Args) >= 3 {
		println("do thing")
	} else {
		bytes, _ := io.ReadAll(os.Stdin)
		println(len(bytes))
	}
}
