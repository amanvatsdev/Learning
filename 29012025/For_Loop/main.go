package main

import (
	"fmt"
)

func EvenOrOdd(Number int) bool {
	A := Number % 2
	if A != 0 {
		return true
	} else {
		return false
	}
}
func DivideBythree(number int) bool {
	Z := number % 3
	if Z == 0 {
		return true
	}
	return false
}

func PrimeNumber(N int) bool {
	if N < 2 {
		return false
	}
	for i := 2; i*i <=N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	// X := 4
	// Sum := 0
	// for Z := 0; Z < X; Z++ {
	// 	fmt.Printf("Value of X =%d \tValue of Sum=%d  Value of Z=%d\n", X, Sum, Z)
	// }

	// for X < 15 {
	// 	fmt.Println("This is the second loop")
	// 	X++
	// }
	for X := 1; X < 16; X++ {
		if X==12{
			fmt.Println("breaking the loop on number 12")
			break
		}
	
		if DivideBythree(X) {
			continue
		}
		fmt.Println(X)
	}

}
