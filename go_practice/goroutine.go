package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i ++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// concurrently runs the goroutine while the normal fxn is running
	go say("world")
	say("hello")
}