package main

import "fmt"

func main() {
	for X := 0; X < 5; X++ {
		fmt.Println("--------------")
		for J := 0; J < 5; J++ {
			fmt.Printf("Value of X is %d \t Value of J is %d\n", X, J)
		}
	}
}
