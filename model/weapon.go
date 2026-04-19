package model

import "math/rand/v2"

type Weapon struct{
	name string
	attackDamage int
}

func SelectRandomWeapon() *Weapon {
	randomIndex := rand.IntN(len(weaponArmory))
	selectedWeapon := weaponArmory[randomIndex]
	return &selectedWeapon
}
