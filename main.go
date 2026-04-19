package main

import (
	"gladiator/core"
	"gladiator/model"
)

func main(){
hero := model.GenerateRandomHero()
adversary := model.GenerateRandomAdversary()

core.Battle(hero, adversary)
}
