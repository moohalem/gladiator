package core

import (
	"fmt"
	"gladiator/internal/gameplay"
	"math/rand/v2"
	"time"
)

// battle now accepts any struct that satisfies the model.Fighter interface
func Battle(p1 gameplay.Fighter, p2 gameplay.Fighter) {
	var turnCounter = 0
	var attacker, defender gameplay.Fighter

	for p1.GetHP() > 0 && p2.GetHP() > 0 {
		// Add a tiny sleep so the terminal output doesn't instantly vanish
		time.Sleep(800 * time.Millisecond)

		if rand.Float64() < 0.5 {
			attacker, defender = p1, p2
		} else {
			attacker, defender = p2, p1
		}

		fmt.Printf("\n--- %s's Turn ---\n", attacker.GetName())

		// Start with base attack damage
		totalDamage := attacker.GetAttackPower()

		// Roll for special attack
		if doSpecialAttack() {
			totalDamage += attacker.SpecialAttack() // Adds special damage!
		} else {
			attacker.Attack() // Just prints standard attack text
		}

		// Calculate final damage mitigated by defense
		actualDamage := totalDamage - defender.GetDefense()
		if actualDamage < 0 {
			actualDamage = 0
		}

		// Apply the damage using our method
		defender.TakeDamage(actualDamage)

		fmt.Printf("%s takes %d damage! ", defender.GetName(), actualDamage)
		turnCounter += 1
		if defender.GetHP() <= 0 {
			fmt.Printf("%s has been defeated!\n", defender.GetName())
			fmt.Printf("The battle took %d turns", turnCounter)
		} else {
			fmt.Printf("%s has %d HP left.\n", defender.GetName(), defender.GetHP())
		}
	}
}

func doSpecialAttack() bool {
	threshold := 0.1
	if rand.Float64() < threshold {
		return true
	}
	return false
}
