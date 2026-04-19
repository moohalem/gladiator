package gameplay

import (
	"fmt"
	"math/rand/v2"
)

type Hero struct{
	Base
	Weapon
	Mana int
}

func GenerateRandomHero() *Hero {
	randomIndex := rand.IntN(len(heroPool))
	selectedHero := heroPool[randomIndex]

	if isWeaponEquip() {
		selectedHero.Weapon = *SelectRandomWeapon()
		fmt.Printf("%s is equipped with the %s weapon!\n", selectedHero.Name, selectedHero.Weapon.name)
		selectedHero.Att += selectedHero.Weapon.attackDamage
		fmt.Printf("The weapon boosted %s's attack %d points.", selectedHero.Name, selectedHero.Weapon.attackDamage)
		return &selectedHero
	}
	fmt.Printf("%s is unarmed!", selectedHero.Name)
	return &selectedHero
}

func isWeaponEquip() bool {
	weaponProbability := 0.1
	randomFloat := rand.Float64()
	if randomFloat < weaponProbability {
		return true
	}
	return false
}

func (h *Hero) SpecialAttack() int {
	if h.Mana == 0 {
		fmt.Printf("Uh oh, %s doesn't have Mana Power.\n", h.GetName())
		return 0
	}
	fmt.Printf("%s casts a MANA SPELL! Power: %d\n", h.Name, h.Mana)
	return h.Mana
}
