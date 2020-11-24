package goroutinechannel

import (
	"fmt"
	"time"
)

func task(channel chan int) {
	time.Sleep(time.Second)
	channel <- 10
	time.Sleep(time.Second)
	channel <- 20
	time.Sleep(time.Second)
	channel <- 30
}

// Execute runs goroutinechannel examples.
func Execute() {
	channel := make(chan int)

	go task(channel)

	fmt.Println(fmt.Sprintf("Read channel value: %d", <-channel)) // Read 10
	fmt.Println(fmt.Sprintf("Read channel value: %d", <-channel)) // Read 20
	fmt.Println(fmt.Sprintf("Read channel value: %d", <-channel)) // Read 30
}
