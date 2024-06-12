package Models

import "gorm.io/gorm"

type Champion struct {
	gorm.Model
	Name    string
	HP      int
	Damage  int
	Agility int

	Item
}

func (ch Champion) GetName() string {
	return ch.Name
}

func (ch Champion) GetHP() int {
	return ch.HP
}
func (ch *Champion) SetHP(add int) {
	ch.HP += add
}

func (ch Champion) GetDamage() int {
	return ch.Damage
}

func (ch *Champion) AddDamage(boostDamage int) {
	ch.Damage += boostDamage
}
func (ch Champion) GetAgility() int {
	return ch.Agility
}
func (ch *Champion) AddAgility(boostAgility int) {
	ch.Agility += boostAgility
}
func (ch Champion) GetItem() Item {
	return ch.Item
}

func (ch *Champion) SetItem(w Item) {
	ch.Item = w
}
