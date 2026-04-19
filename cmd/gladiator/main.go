package main

import (
	core "gladiator/internal/engine"
	"gladiator/internal/gameplay"
)

func main(){
	hero := gameplay.GenerateRandomHero()
	adversary := gameplay.GenerateRandomAdversary()

	core.Battle(hero, adversary)
}
