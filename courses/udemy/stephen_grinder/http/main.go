package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// main function will execute a request to an URL an print the response
func main() {
	// Built-in http lib receives the response and an error if needed
	resp, err := http.Get("https://www.google.com")

	// Here is the error handling
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		// Building the slice of byte which will contain the response body. A requirement of the Read interface
		// bs := make([]byte, 99999)
		// resp.Body.Read(bs)
		// fmt.Println(string(bs))

		lw := logWriter{}
		// Here is a more straightforward approach, using the Writer interface of Stdout.
		// The Copy function takes data from a reader and directly copies that data into the writer
		//  io.Copy(os.Stdout, resp.Body)

		// Here is the custom writer
		io.Copy(lw, resp.Body)
	}
}

// Let's implement our own custom writer, because the f* the police
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println(len(bs), "bytes written")
	fmt.Println()
	return len(bs), nil
}
