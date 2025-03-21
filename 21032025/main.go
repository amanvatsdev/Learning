package main

import "fmt"

func Name() {
	fmt.Println("My Name is Aman")
} //with 0 Parameter

func Bull(name string) {
	fmt.Println("My Name is ", name)
} // func with 1 parameter

func Raghu(name string) string {
	return fmt.Sprintf("My Father Name is %s", name)
} //1 parameter and 1 return

func Raju(a int, b string) (string, int) {
	return fmt.Sprint(b, " is the name and the age is"), a
}

func main() {
	Name()
	Bull("Aman Vats")
	A := Raghu("Naresh Kumar")
	fmt.Println(A)
	D, E := Raju(21, "Aman")
	fmt.Println(D, E)
}
