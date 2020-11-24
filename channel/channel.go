package channel

import "fmt"

// Execute runs channel examples.
func Execute() {
	channel := make(chan int, 1)

	channel <- 1 // Send channel data (write)
	<-channel    // Receive channel data (read)

	// See the example below:
	channel <- 2                                                      // Send the value (2) for the channel
	fmt.Println(fmt.Sprintf("Received channel value: %d", <-channel)) // Read channel sent value (2)
}
