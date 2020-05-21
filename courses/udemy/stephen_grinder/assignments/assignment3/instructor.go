package main

import (
	"fmt"
	"io"
	"os"
)

// This is the course instructor solution, not as direct.
func mainly() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
