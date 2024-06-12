package Models

type Invoker struct {
	Stats
	//Weapon
	Name string
}

func (hero *Invoker) GetName() string {
	hero.Name = "Invoker"
	return hero.Name
}

/*
	func (hero Invoker) GetWeapon() Weapon {
		return hero.Weapon
	}
*/
func (hero Invoker) GetHP() int {
	return hero.HP
}
func (hero *Invoker) SetHP(hpChange int) {
	hero.HP += hpChange
}

func (hero Invoker) GetAgility() int {
	return hero.Agility
}

func (hero *Invoker) SetAgility(agilityChange int) {
	hero.Agility += agilityChange
}
func (hero Invoker) GetDamage() int {
	return hero.Damage
}

func (hero *Invoker) SetDamage(damageChange int) {
	hero.Damage += damageChange
}
