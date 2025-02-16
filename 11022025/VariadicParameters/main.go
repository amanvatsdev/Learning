package main

import "fmt"

func main() {
	fmt.Println("hi")

	fmt.Println(One(1,2,3,4,5,6,7,8,9))

	M:=[]int{5,2,3,4,5,6,4,5,}

	Two("The slice value ",M...)
}

func One(X ...int) (int ,string){
	Sum := 0

	for _, V := range X {
		Sum += V
	}
	return Sum ,"This is the Output"
}

func Two(Y string,X ...int){
	// fmt.Println(Y,X)
	for _,Value:=range X{
		fmt.Println(Y,Value)
	}
}
