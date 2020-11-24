package blockingchannel

import (
	"fmt"
	"time"
)

func task(channel chan int) {
	time.Sleep(time.Second)
	channel <- 1 // Blocking operation
	fmt.Println("Channel value read!")
	time.Sleep(5 * time.Second)
	fmt.Println("After sleep")
	channel <- 10 // Blocking operation
}

// Execute runs blockingchannel examples.
func Execute() {
	channel := make(chan int)

	go task(channel)

	fmt.Println(fmt.Sprintf("Channel value: %d", <-channel)) // Blocking operation
	fmt.Println("After read channel value")
	fmt.Println(fmt.Sprintf("Channel value: %d", <-channel)) // Blocking operation
}
