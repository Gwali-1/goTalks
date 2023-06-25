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
	go increaseCounter(int)
	wg.Add(1)
	go increaseCounter(int)
	wg.Wait()
	log.Printf("FINAL COUNTER VALUE: %v", COUNTER)
}

func increaseCounter(iter int) {
	defer wg.Done()
	for i := 0; i < iter; i++ {
		COUNTER++
	}
}
