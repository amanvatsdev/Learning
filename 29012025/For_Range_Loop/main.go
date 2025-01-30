package main

import "fmt"

func main() {
	var Name = []string{"Aman", "Sinu", "Vikash"}

	for Index, Value := range Name {
		fmt.Printf("%d. Name is %s\n", Index+1, Value)
	}
Z:=map[string]int{
	"Aman"  : 22,
	"Sinu": 25,
	"Vikash":32,
}
for Index,Value:=range Z{
	fmt.Printf("Name is %s \t Age is %d\n",Index,Value)
}
}
