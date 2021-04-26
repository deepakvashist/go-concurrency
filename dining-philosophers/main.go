package main

import (
	"fmt"
	"sync"
	"time"
)

var eatWgroup sync.WaitGroup

func say(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}

type fork struct {
	mutex sync.Mutex
}

type philosopher struct {
	id        int
	leftFork  *fork
	rightFork *fork
}

func (p philosopher) eat() {
	for x := 0; x < 3; x++ {
		p.leftFork.mutex.Lock()
		p.rightFork.mutex.Lock()

		say("eating", p.id)
		time.Sleep(time.Second)

		p.leftFork.mutex.Unlock()
		p.rightFork.mutex.Unlock()

		say("finished eating", p.id)
		time.Sleep(time.Second)
	}

	eatWgroup.Done()
}

func main() {
	count := 5

	// Create forks
	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id:        i,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%count],
		}

		eatWgroup.Add(1)

		go philosophers[i].eat()
	}

	eatWgroup.Wait()
}
