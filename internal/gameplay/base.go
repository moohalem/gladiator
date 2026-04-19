package gameplay

import "fmt"

type Fighter interface {
	GetName() string
	GetHP() int
	TakeDamage(amount int)
	GetAttackPower() int
	GetDefense() int
	Attack()
	SpecialAttack() int // We make this return an int so we know how much damage to add!
}

type Base struct {
	Name string
	HP int
	Att int
	Def int
}

func (b *Base) GetName() string  { return b.Name }
func (b *Base) GetHP() int       { return b.HP }
func (b *Base) GetAttackPower() int { return b.Att }
func (b *Base) GetDefense() int     { return b.Def }

func (b *Base) TakeDamage(amount int) {
	b.HP -= amount
}

func (b *Base) Talk(){
	fmt.Printf("%s is slamming hands all around. It is very angry and wants to fight you!\n", b.Name)
}
func (b *Base) Attack() {
	fmt.Printf("%s attacks for %d damage\n", b.Name, b.Att)
}
