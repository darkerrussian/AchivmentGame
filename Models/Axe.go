package Models

import (
	"math/rand"
	"time"
)

type Axe struct {
	Damage int
	Name   string
}

func (axe Axe) AgilityDebuff() int {
	return -2
}

func (axe Axe) GetName() string {
	axe.Name = "Axe"
	return axe.Name
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

func (axe Axe) GetDamage() int {

	axe.Damage = rand.Intn(4) + 7
	return axe.Damage

}

func (axe Axe) GetUnique(hero Hero) {

	hero.SetAgility(axe.AgilityDebuff())

}
