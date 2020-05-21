package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type logWriter struct{}

func main() {
	// Get the command line arguments. Beware! This will take everithing in the command, even PATH/main
	args := os.Args

	// So instead of recovering the first element of the slice take the second, but this will fail if the program is built
	// There's one parameter at least
	if len(args) > 1 && args[1] != "" {
		fn := args[1]
		text, err := ioutil.ReadFile(fn)

		if err != nil {
			panic(err)
		}
		fmt.Println(string(text))
	} else {
		fmt.Println("Nothing to see")
	}
}
