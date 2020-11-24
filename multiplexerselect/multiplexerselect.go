package multiplexerselect

import (
	"fmt"
	"time"
)

func talk(people string) <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			channel <- fmt.Sprintf("%s talking %d", people, i)
		}
	}()

	return channel
}

func merge(channelA, channelB <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case s := <-channelA:
				channel <- s
			case s := <-channelB:
				channel <- s
			}
		}
	}()

	return channel
}

// Execute runs multiplexerselect examples.
func Execute() {
	channel := merge(
		talk("Deepak"),
		talk("Maria"),
	)

	fmt.Println(<-channel, <-channel)
	fmt.Println(<-channel, <-channel)
	fmt.Println(<-channel, <-channel)
}
