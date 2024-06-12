package Models

type Weapon interface {
	GetDamage() int
	GetUnique(hero Hero)
	GetName() string
}
