package model

import "fmt"

type Base struct {
	Name string
	HP int
	Att int
	Def int
}

func (b *Base) GetName() string { return b.Name }
func (b *Base) GetHP() int      { return b.HP }
func (b *Base) GetAtt() int     { return b.Att }
func (b *Base) GetDef() int     { return b.Def }

func (b *Base) TakeDamage(amount int) {
	b.HP -= amount
}

func (b *Base) Talk(){
	fmt.Printf("%s is slamming hands all around. It is very angry and wants to fight you!\n", b.Name)
}
func (b *Base) Attack() {
	fmt.Printf("%s attacks for %d damage\n", b.Name, b.Att)
}
