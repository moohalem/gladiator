package gameplay

import (
	"fmt"
	"math/rand/v2"
)

type Adversary struct{
	Base
	SpcAtt int
}

func GenerateRandomAdversary() *Adversary{
	randomIndex := rand.IntN(len(adversaryPool))
	selectedAdversary := adversaryPool[randomIndex]
	return &selectedAdversary
}

func (a *Adversary) SpecialAttack() int {
	fmt.Printf("%s uses SPECIAL ATTACK! Power: %d\n", a.Name, a.SpcAtt)
	return a.SpcAtt
}
