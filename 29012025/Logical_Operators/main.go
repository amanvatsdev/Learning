// In this i will learn about how to work with logical operators
// Logical operators
// 1.&&  (And)
// 2.|| (Or)
// 3. ! (Not Equal to )
package main

import "fmt"

func main() {
	X := 46

	Y := 30
	if X < 45 || X == 45 {
		fmt.Println("Either X is less than 45 or equal to 45")
	} else {
		fmt.Println("X is greater than 45")
	}
	if X < Y && X < 45 {
		fmt.Println("X is less than Y as well as less than 45")
	}
	if X != Y || X > Y {
		fmt.Println("Either X is Greater than Y or X is not equal to Y")
	} else {
		fmt.Println("X is equal to Y")
	}
}
