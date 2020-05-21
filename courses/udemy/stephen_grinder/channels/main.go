package main

import (
	"fmt"
	"net/http"
	"time"
)

// This works indeed, but there is no concurrence, and every request hast to wait the previous being resolved
func main() {
	links := []string{
		"https://google.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://amazon.com",
	}

	// Add a channel to handle execution of go routines
	c := make(chan string)

	for _, link := range links {

		// go checkLink(link) // This does not work by itself, nothing is printed. May be reaches the end of file before the first request is resolved?
		go checkLink(link, c)
	}
	// Wait to channel to write as many times as links are
	//for i := 0; i < len(links); i++ {
	// Commented to step forward
	//fmt.Println(<-c)

	// Instead to run just once, it is going to check this continuosly with an infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Other way to ping almost constantly infinitely
	for l := range c {
		// Using a function literal (lambda or anonymous)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

// Code commented below is the non-optimized approach
// func checkLink(link string) {
func checkLink(link string, c chan string) { // Add a channel to handle the execution
	_, err := http.Get(link) // Here the execution is frozen. No doubt this can be improved

	if err != nil {
		fmt.Println(link, "seems down")
		c <- link
		return
	}
	fmt.Println(link, "Status OK")
	c <- link

}
