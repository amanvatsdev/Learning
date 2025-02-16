// in this we will take function as an parameter/argument
package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}

func Calculator(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func main() {
	X := Calculator(90,45, subtract)
	Y := Calculator(100,54, add)
	fmt.Println("subtraction", X)
	fmt.Println("Add", Y)
	fmt.Printf("type of X is :%T\n", X)
	fmt.Printf("type of Y is :%T\n", Y)
	fmt.Printf("type of add is :%T\n", add)
	fmt.Printf("type of subtract is :%T\n", subtract)
}
