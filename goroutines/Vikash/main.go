package main

import (
	"fmt"
	// "sync"
)

func main() {

	go new()
	new2()
}

func new() {
	for i := 0; i < 50; i++ {
		fmt.Println("go routine 2")
	}
}

func new2() {
	for i := 0; i < 10; i++ {
		fmt.Println("go routine 1")
	}
}
