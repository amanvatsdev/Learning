package main

import "fmt"

func sum() func() int {
	Sum := 0
	return func() int {
		Sum += 9
		return Sum
	}
}

func main() {

	fmt.Printf(" Sum:%T\n ", sum())

	X := sum()
	fmt.Println(X()) // 09
	fmt.Println(X()) // 18
	fmt.Println(X()) // 27
	fmt.Println(X()) // 36
	fmt.Println(X()) // 45
	fmt.Println(X()) // 54
	fmt.Println(X()) // 63
	fmt.Println(X()) // 72
	fmt.Println(X()) // 81
	fmt.Println(X()) // 90
}
