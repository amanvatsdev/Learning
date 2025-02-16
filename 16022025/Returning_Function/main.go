// in This we will learn about Returning a function
package main

import "fmt"

// Simpple Function
func Simple() int {
	return 456
}

func Function_caller() func() int {

	return func() int {
		return 45
		}
	}



func main() {
	Caller1 := Simple()
	Caller2 := Function_caller()

	fmt.Println(Caller1)
	fmt.Println(Caller2())

	fmt.Printf("Simple:%T\n",Simple())
	fmt.Printf("Function Caller:%T\n",Function_caller())
	fmt.Printf("Caller1:%T\n",Caller1)
	fmt.Printf("Caller2:%T\n",Caller2)
	fmt.Printf("Caller2():%T\n",Caller2())
}
