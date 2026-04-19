package model

type Fighter interface {
	GetName() string
	GetHP() int
	TakeDamage(amount int)
	GetAtt() int
	GetDef() int
	Attack()
	SpecialAttack() int // We make this return an int so we know how much damage to add!
}
