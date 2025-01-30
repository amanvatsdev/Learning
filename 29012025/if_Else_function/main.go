// In this we practiced if , else , if else
package main

import "fmt"

const Y = 40

func main() {
	fmt.Println("Program Started")
	X := 45

	if X < Y {
		fmt.Println("X is less than the constant")
	}
	if X < Y {
		fmt.Println("The Value Of X is less than Constant Value")
	} else {
		fmt.Println("The Value of X is Not less than the Constant ")
	}
	if X < Y {
		fmt.Println("The Value of X is less than 45 which is the value of constant")
	} else if X == Y {
		fmt.Println("The Value is equal to constant")
	} else {
		fmt.Println("X is Not less  than Constant Y")
	}
}
