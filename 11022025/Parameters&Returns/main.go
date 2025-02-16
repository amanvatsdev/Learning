/*       syntax=func(Receiver)Identifier(parameter)return{Code}
 
In this we will practice functions with different condition
1. no paramete no return
2. 1 param 0 return
3. 1 param 1 return
4. 2 param 2 return
*/

package main

import "fmt"

func main() {
	One()
	Two("Aman")
	fmt.Println(Three("Aman Vats"))
	fmt.Println(Four("Aman","Vats"))

}

func One() {
	fmt.Println("HI World")
}

func Two(A string) {
	fmt.Println("Hi World My name is", A)
}

func Three(A string) string {
	return fmt.Sprint("My Full Name is ", A)
}

func Four (X,Y string)string{
	return fmt.Sprint("My first Name is ",X," and last name is ",Y)
}
