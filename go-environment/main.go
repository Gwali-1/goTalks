package main

import (
	"fmt"
	"time"
	"flag"
)

var (
	name = flag.String("name","john", "what is your name")
)

func main() {
	fmt.Println(time.Now().Day())
	flag.Parse()
	fmt.Println(*name)

}
