package Models

type Pudge struct {
	Stats
	//Weapon
	Name string
}

func (hero *Pudge) GetName() string {
	hero.Name = "Pudge"
	return hero.Name
}

/*
	func (hero *Pudge) GetWeapon() Weapon {
		return hero.Weapon
	}
*/
func (hero Pudge) GetHP() int {
	return hero.HP
}
func (hero *Pudge) SetHP(hpChange int) {
	hero.HP += hpChange
}

func (hero Pudge) GetAgility() int {
	return hero.Agility
}

func (hero *Pudge) SetAgility(agilityChange int) {
	hero.Agility += agilityChange
}
func (hero Pudge) GetDamage() int {
	return hero.Damage
}

func (hero *Pudge) SetDamage(damageChange int) {
	hero.Damage += damageChange
}
