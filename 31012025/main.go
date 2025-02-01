package main

import "fmt"

func main() {
	//#95
	// Var1 := [5]string{"Aman", "Sam", "Ravi", "Ajay"}
	// fmt.Printf("Index Number is %d\n and th value is %s\n", len(Var1), Var1)
	// Var1[4] = "Vikash"
	// fmt.Printf("Index Number is %d\n and th value is %s\n", len(Var1), Var1)
	// Var2 := [...]string{"Gulshaan", "Pappi"}
	// fmt.Printf("Index Number is %d\n and th value is %s\n", len(Var2), Var2)

	//#97
	// Var1 := []int{45, 24, 265, 65, 95}

	// fmt.Println(Var1)

	// for _, Value := range Var1 {
	// 	fmt.Printf("Value of V1 is : %d\n", Value)
	// }
	// fmt.Println(Var1[0])
	// fmt.Println(Var1[1])
	// fmt.Println(Var1[2])
	// fmt.Println(Var1[3])
	// fmt.Println(Var1[4])

	// fmt.Println("lenth of Var1 is ", len(Var1))

	// // #98
	// Var1 = append(Var1, 45, 65)
	// fmt.Printf("The len of Var1 is :%d\n The Value of Var1:%d\n", len(Var1), Var1)

	// //99 slicing a slice
	// Var1 := []int{45, 24, 265, 65, 95, 65, 78}

	// fmt.Println("Value of Var1 is ", Var1)
	// fmt.Println("----------------------------------")

	// fmt.Printf("Showing its type and structure of Var1:--%#v\n ", Var1)

	// //Slicing  a slice
	// //Inclusive:Exclusive
	// fmt.Printf("[Inclusive:Exclusive] : %#v\n", Var1[1:4])
	// fmt.Println("--------------------------------------------------------------")
	// //Inclusive:
	// fmt.Printf("[Inclusive :] %#v\n", Var1[3:])
	// fmt.Println("--------------------------------------------------------------")
	// //:Exclusive
	// fmt.Printf("[:Exclusive]: %#v\n", Var1[:6])
	// fmt.Println("--------------------------------------------------------------")
	// //[:]
	// fmt.Printf("[:] : %#v\n", Var1[:])
	// fmt.Println("--------------------------------------------------------------")

	//100 Deleting a slice
// 	Ages := []int{52, 64, 95, 78, 456, 65, 55}
// fmt.Println("----------------------------------------------------------------------------")
// fmt.Printf("%#v\n", Ages)

// Ages = append(Ages[:3], Ages[5:]...) //4 and 5 will be deleted by this
// fmt.Printf("%#v\n", Ages)

// fmt.Println("----------------------------------------------------------------------------")
// Date := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// fmt.Printf("%#v\n", Date)
// Date = Date[1:] // this will delete the first index number
// fmt.Printf("%#v\n", Date)

// fmt.Println("----------------------------------------------------------------------------")
// 	Day := []int{45, 6, 4, 84, 5, 64, 5, 45}
// 	fmt.Printf("%#v\n", Day)
// 	Day = Day[:len(Day)-1]
// 	fmt.Printf("%#v\n", Day)




//#101 Practicing make with slice
// Names:=make([]string,0,10)
// fmt.Printf("The lenth is %d and Capacity is %d and the data is %v\n",len(Names),cap(Names),Names)
// Names=append(Names, "aman","Raj","Pawan","Deleep","Raghu","Niyat","bhuwan","Vikash","Raghu","Niyat","bhuwan")
// fmt.Printf("The lenth is %d and Capacity is %d and the data is %v",len(Names),cap(Names),Names)



//#102 Multi dimentional slice
Names:=[]string{"Aman","Raj","Pawan","Deleep"}
LastNames:=[]string{"Vats","Gupta","kumar","yadav"}

FullName:=[][]string{Names,LastNames}
fmt.Println(FullName)



}
