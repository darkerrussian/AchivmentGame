package Models

import "gorm.io/gorm"

type Champion struct {
	gorm.Model
	Name     string
	HP       int
	Damage   int
	Agility  int
	Killable bool
	Kills    int
}

func (ch Champion) GetKills() int {
	return ch.Kills
}

func (ch *Champion) AddKill() int {
	return ch.Kills + 1
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
