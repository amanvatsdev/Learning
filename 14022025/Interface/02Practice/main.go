//go:generate stringer -type=Day
//go:generate stringer -type=Day
package main

import "fmt"

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)


func (d Day) String() string {
	return []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

func main() {
	fmt.Println(Wednesday) // Output: Wednesday
}





























// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// type Person struct {
// 	Name string
// }

// func (p Person) String() string {
// 	return fmt.Sprint("The Name is ",p.Name)
// }

// type Age int

// func (a Age)String()string{
// 	return fmt.Sprint("The Age is ",strconv.Itoa(int(a)))
// }

// func main(){
// 	X:=Person{"Aman"}

// 	Y:=22

// 	fmt.Println(X)
// 	fmt.Println(Y)
// }