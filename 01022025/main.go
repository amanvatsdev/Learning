package main

import (
	"fmt"
)

func main() {

	// #106
	Var1 := []int{45, 46, 85, 95, 74, 65, 45}

	Var2 := make([]int, 10)

	fmt.Printf("VAr1 Lenth=%d Capacity =%d,Value=%d\n", len(Var1), cap(Var1), Var1)
	fmt.Printf("VAr2 Lenth=%d Capacity =%d,Value=%d\n", len(Var2), cap(Var2), Var2)

	copy(Var2, Var1)

	fmt.Printf("VAr1 Lenth=%d Capacity =%d,Value=%d\n", len(Var1), cap(Var1), Var1)
	fmt.Printf("VAr2 Lenth=%d Capacity =%d,Value=%d\n", len(Var2), cap(Var2), Var2)
	fmt.Println("------------------------------------------------------------------------------")
	Var1[1] = 9999999

	fmt.Printf("VAr1 Lenth=%d Capacity =%d,Value=%d\n", len(Var1), cap(Var1), Var1)``
	fmt.Printf("VAr2 Lenth=%d Capacity =%d,Value=%d\n", len(Var2), cap(Var2), Var2)

	Var1[1] = 9999

	fmt.Printf("VAr1 Lenth=%d Capacity =%d,Value=%d\n", len(Var1), cap(Var1), Var1)
	fmt.Printf("VAr2 Lenth=%d Capacity =%d,Value=%d\n", len(Var2), cap(Var2), Var2)

}
