package main

import "time"
import "fmt"

func main() {

	c := boringGen("mesage")
	p := boringGen("ffreestyle me")

	x := fanInWithSelect(c, p)

	for i := 0; i < 5; i++ {
		fmt.Println(<-x)
	}

}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%v %v", msg, i)
		time.Sleep(time.Second)
	}
}

// a fucntion that takes a channel and launches goroutines to recieve value into another chanel and returns that channel
// we can then read from that channel , this prevents the issue of having the channel read being snchronous thus reading from one
// channel while the other waits even if one is faster
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
func fanInWithSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}


// function that that returns a channel, lauches a go routine that feeds data into that channel and we call the function get the
// the channel back and read from it
func boringGen(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%v %v", msg, i)
			time.Sleep(time.Second)

		}
	}()

	return c
}
