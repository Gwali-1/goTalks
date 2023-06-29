package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	name = flag.String("name", "john", "what is your name")
)

func main() {
	fmt.Println(time.Now().Day())
	flag.Parse()
	fmt.Println(*name)

	//getting file info
	fileInfo, err := os.Stat("new.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}

fmt.Println(fileInfo.Name())
fmt.Println(fileInfo.ModTime())
fmt.Println(fileInfo.IsDir())




}
