package main

import "fmt"

type person struct {
	name string
}

type SecretAgent struct {
	person
	SecretAgent bool
}

type human interface{
	Voice()
}
func main() {
	P1:=person{
		name: "Aman",
	}
	fmt.Println(P1.name)


	P2:=SecretAgent{
		person: person{name: "Raghu"},
		SecretAgent: true,
	}
	fmt.Println(P2)
	
	P1.Voice()
	P2.Voice()

	LoudVoice(P1)
	LoudVoice(P2)
	
}

func (p person) Voice() {
	fmt.Println("My Name is ", p.name)
}
func (Sc SecretAgent)Voice(){
	fmt.Println("The Name of Secret Agent is ",Sc.name)
}
func LoudVoice(h human){
	h.Voice()
}

//101159536042
774359722774
7056486110