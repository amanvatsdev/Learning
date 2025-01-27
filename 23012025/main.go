package main

import (
	"fmt"

	"github.com/amanvatsdev/Learning/23012025/TagVersion"
	"github.com/amanvatsdev/Learning/23012025/dependecy/Bark"
	"github.com/amanvatsdev/Learning/23012025/dependecy/Sound"
)

func Tag14() {
	fmt.Println("This is for tag 1.1.1")
}

func main() {
	AdultDog := Bark.AuldtDog()
	Puupy := Bark.Puppy()
	CatSound := Sound.Cat()
	HumanSound := Sound.Human()
	AnimalSounds := []string{AdultDog, Puupy, CatSound, HumanSound}

	for SeraialNumber, Value := range AnimalSounds {

		fmt.Println(SeraialNumber+1, Value)

		//using of tag version package here
		TagVersion.PracticingVersion()
		Tag14()


	// this is brought from the Hands on Exercise Directory 
	
	}
	
}
