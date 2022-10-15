package main

import (
	"fmt"
	"time"
)

func main() {
	namesChan := make(chan string, 1)

	go readNames(namesChan)

	names := []string{"Deepak", "João", "Maria", "Foo", "Bar"}
	for _, name := range names {
		fmt.Println("writing name: ", name)
		namesChan <- name
	}

	time.Sleep(20 * time.Second)
}

func readNames(namesChan chan string) {
	fmt.Println("start reading names:")

	for name := range namesChan {
		fmt.Println("name -> ", name)
		time.Sleep(3 * time.Second)
	}
}
