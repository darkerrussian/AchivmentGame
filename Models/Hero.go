package Models

type Hero interface {
	GetName() string
	GetHP() int
	SetHP(hpChange int)
	GetAgility() int
	SetAgility(agilityChange int)
	GetDamage() int
	SetDamage(damageChange int)
	//GetWeapon() Weapon
}

type Stats struct {
	HP      int
	Agility int
	Damage  int
}
