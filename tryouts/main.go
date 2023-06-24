package main

import (
	"log"
	"os"
	"strconv"
	"sync"
)

var COUNTER int = 0
var iterations string = os.Args[1]
var wg sync.WaitGroup

func main() {

	int, err := strconv.Atoi(iterations)
	if err != nil {
		log.Panic("Pass in iteger number of iterations")
	}

	log.Printf("INITIAL COUNTER VALUE: %v", COUNTER)
	wg.Add(1)
	go increaseCounter(int, "worker 1")
	wg.Add(1)
	go increaseCounter(int, "worker 2")
	wg.Wait()
	log.Printf("FINAL COUNTER VALUE: %v", COUNTER)
}

func increaseCounter(iter int, name string) {
	defer wg.Done()
	for i := 0; i < iter; i++ {
		COUNTER++
	}
}
