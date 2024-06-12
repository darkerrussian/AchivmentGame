package Models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name   string
	Damage int
}

func (i Item) GetUnique(ch *Champion) {
	if ch.GetName() == "Axe" {
		ch.AddAgility(-2)
	}
}

func (i Item) GetName() string {
	return i.Name
}
func (i Item) GetDamage() int {
	return i.Damage
}
