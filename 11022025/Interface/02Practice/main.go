/*
1.Define an inerface 'Animal'with a method 'Sound() string'.
2.Create Structs 'Dog' ,'Cat', and 'Cow'.
3.Implement 'Sound()' for each struct to return thier respective sounds :

	-Dog:"Woof Woof"
	-Car: "Meow"
	-Cow:"Moo"

4.Write a function 'makeSound(Animal)' to print the sounds.
5.Call 'makeSound()' for each animal in the 'main()'function.
*/
package main

import "fmt"


type Animal interface{
	Sound()string
}

type Dog struct{}
func (d Dog)Sound()string{
	return "Woof Woof"
}

type Cat struct{}
func (c Cat)Sound()string{
	return "Meow!"
}

type Cow struct{}
func (cw Cow)Sound()string{
	return "Mooo!"
}

func MakeSoundAnimal (A Animal){
	fmt.Println(A.Sound())
}
func main(){
Dog:=Dog{}
Cat:=Cat{}
Cow:=Cow{}

MakeSoundAnimal(Dog)
MakeSoundAnimal(Cat)
MakeSoundAnimal(Cow)
}
