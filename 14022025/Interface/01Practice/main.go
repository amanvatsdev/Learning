package main

import "fmt"

type person struct {
	Name string
}

type SecretAgent struct {
	person
}

func (p person) speak() {
	fmt.Println(p.Name,"is a simple Person")
}

func (sa SecretAgent)speak(){
	fmt.Println(sa.Name,"is a Secret Agent")
}

type human interface{
	speak()
}

func main() {
	P1:=person{"Aman Vats"}

	P2:=SecretAgent{person: person{"Raghu"}}

	P1.speak()
	P2.speak()
}