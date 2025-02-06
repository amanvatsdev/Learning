package main

import "fmt"

func main() {
	//#117 Range loop over map

	// Name := map[int]string{
	// 	1: "Aman",
	// 	2: "Raju",
	// 	4: "Pankaj",
	// }
	// for Key, Value := range Name {
	// 	fmt.Printf("%d:%s\n",Key,Value)
	// }

	//#118 delete a slice and see the results
	// Name := map[int]string{
	// 	1: "Aman",
	// 	2: "Raju",
	// 	3: "Pankaj",
	// }
	// fmt.Println(Name)
	// delete(Name,2)
	// fmt.Println(Name)

	//#119 Use of Ok in map with if statement
	X := map[int]string{
		1: "Aman",
		2: "Vikash",
		3: "Sinu",
	}
	fmt.Println("1", X)
	delete(X, 2)
	fmt.Println("2", X)
	X[2]="Vikash Parashar"
	fmt.Println("3", X)
	
	if Value, ok := X[2];ok {
		fmt.Println(Value)
	} else {
		fmt.Println("Value not found of variable ")
	}

}
