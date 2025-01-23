package main

import (
	"fmt"

	"github.com/amanvatsdev/Learning/23012025/dependecy/Bark"
	"github.com/amanvatsdev/Learning/23012025/dependecy/Sound"
)

func main() {
	AdultDog:=Bark.AuldtDog()
	Puupy:=Bark.Puppy()
	CatSound:=Sound.Cat()
	HumanSound:=Sound.Human()
	AnimalSounds:=[]string{_,AdultDog,Puupy,CatSound,HumanSound}
	for SeraialNumber,Value:=range AnimalSounds{
		fmt.Println(SeraialNumber,Value)
	}
}