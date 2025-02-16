package main

import "fmt"

type Student struct {
	Name  string
	Marks []int
}

func (P Student) Result() (bool, float64) {
	Sum := 0
	Subjects := 0
	for _, V := range P.Marks {
		Sum += V
		Subjects++
	}
	Percent := float64(Sum) / float64(Subjects)
	return Percent >= 40, Percent
}

func main() {
	S1 := Student{
		Name:  "Aman",
		Marks: []int{25, 30, 100, 28, 26},
	}
	Passed, Percent := S1.Result()

	if Passed {
		fmt.Println(S1.Name, "Has passed with percentage", Percent)
	} else {
		fmt.Println(S1.Name, "Has Failed with Percentage", Percent)
	}

}

// type Rectangle struct{
// 	Lenth int
// 	Width int
// }
// func (B Rectangle)Area()int{
// 	Area:=B.Lenth*B.Width

// 	return Area
// }
// func main(){
// 	R1:=Rectangle{
// 		Lenth: 10,
// 		Width: 5,
// 	}
// 	fmt.Println("Area Of Rectangle:",R1.Area())
// }











// type Person struct{
// 	Name string
// }
// func (B Person)Greet(message string){
// 	fmt.Println(message,B.Name,",Welcome!")
// 	fmt.Println("Good Morning ",B.Name,"!")
// }

// func main(){
// 	P1:=Person{Name: "Aman"}

// 	P1.Greet("Hello")
// }
















// type Book struct{
// 	Title string
// }

// func main(){
// B1:=Book{Title: "The Go Programming Language"}

// B1.ShowTitle()
// }

// func (B Book)ShowTitle(){
// 	fmt.Println("Book Title:",B.Title)
// }
