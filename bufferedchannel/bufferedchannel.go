package bufferedchannel

import "fmt"

func task(channel chan int) {
	channel <- 0
	fmt.Println("Executed - 0")
	channel <- 1
	fmt.Println("Executed - 1")
	channel <- 2
	fmt.Println("Executed - 2")
	channel <- 3
	fmt.Println("Executed - 3")
	channel <- 4
	fmt.Println("Executed - 4")
	channel <- 5
	fmt.Println("Executed - 5")
	channel <- 6
	fmt.Println("Executed - 6")
	channel <- 7
	fmt.Println("Executed - 7")
	channel <- 8
	fmt.Println("Executed - 8")
	channel <- 9
	fmt.Println("Executed - 9")
}

// Execute runs bufferedchannel examples.
func Execute() {
	channel := make(chan int, 3)

	go task(channel)

	fmt.Println(fmt.Sprintf("Read channel value: %d", <-channel))
}
