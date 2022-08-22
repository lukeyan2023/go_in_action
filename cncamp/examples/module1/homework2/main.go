package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	messages := make(chan int, 10)

	// consumer
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for _ = range ticker.C {
			fmt.Printf("consumer print: %d\n", <-messages)
		}
	}()

	// producer
	for _ = range ticker.C {
		messages <- rand.Intn(50)
	}

}
