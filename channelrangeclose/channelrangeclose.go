package channelrangeclose

import (
	"fmt"
	"time"
)

func isPrimeNumber(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func primes(number int, channel chan int) {
	start := 2
	for i := 0; i < number; i++ {
		for prime := start; ; prime++ {
			if isPrimeNumber(prime) {
				channel <- prime
				start = prime + 1
				time.Sleep(time.Millisecond * 200)
				break
			}
		}
	}

	close(channel)
}

// Execute runs channelrangeclose examples.
func Execute() {
	channel := make(chan int, 30)

	// cap(channel) is the channel capacity.
	go primes(cap(channel), channel)

	for prime := range channel {
		fmt.Println(fmt.Sprintf("Prime number: %d", prime))
	}

	fmt.Println("End of program!!")
}
