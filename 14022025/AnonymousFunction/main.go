package main

import "fmt"

func main() {
	x:=func (){
		fmt.Println("My Name is Aman")
	}
	
	name:=func(name string) {
		fmt.Println("This is my Name",name)
	}
	Z:=func (x,y int){
		Sum:=x+y
		fmt.Println(Sum)
	}
	
	x()
	name("Aman")
	Z(45,65)
}