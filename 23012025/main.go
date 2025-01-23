package main

import (
	"fmt"

	"github.com/amanvatsdev/Learning/23012025/TagVersion"
	"github.com/amanvatsdev/Learning/23012025/dependecy/Bark"
	"github.com/amanvatsdev/Learning/23012025/dependecy/Sound"
)

func main() {
	AdultDog:=Bark.AuldtDog()
	Puupy:=Bark.Puppy()
	CatSound:=Sound.Cat()
	HumanSound:=Sound.Human()
	AnimalSounds:=[]string{AdultDog,Puupy,CatSound,HumanSound}
	
	for SeraialNumber,Value:=range AnimalSounds{
	
		fmt.Println(SeraialNumber+1,Value)

	//using of tag version package here
	TagVersion.PracticingVersion()
	}
}