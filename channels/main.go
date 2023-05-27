package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}
	for l := range c { // this is an alternative syntax to the obvious infinite for loop previously written
		//	time.Sleep(5 * time.Second) putting this pause function here causes some problem, think about it,refer to video no 80

		go func() { // here we have used a function literal
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()

	}

	// this is also a blocking code, the main go routine waits here till something comes into the channel and as soon as a signal comes into channel the main routine termainates and it doesn't wait for other goroutine to finish
	// here there is a slight delay in printing of site status between different sites, so we should come up with a solution to make all the request at the same time, so where go routines comes into picture
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Get function here is creating a blocking call

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}

// we write go keyword only is written in front of function call
// channels are used so that main goroutine is aware that it's child goroutines are not yet completed
// channels are used to communicate between different routines
// we can create
