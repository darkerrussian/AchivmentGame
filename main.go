package main

import (
	"AchivmentGame/Database"
	"AchivmentGame/Models"
	"AchivmentGame/Server"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	db, err := Database.InitDB()
	http.HandleFunc("/", Server.Handle)
	//http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Models.Champion{})
	db.AutoMigrate(&Models.Item{})
	//item :=Models.Item{Name: "Axe", Damage: 7}
	//db.Create(&item)
	var item Models.Item
	db.First(&item, 2) //Axe
	var championI Models.Champion
	db.First(&championI, 7) //invoker
	championI.Item = item
	var championP Models.Champion
	db.First(&championP, 2) //Pudge
	championP.Item = item
	//fmt.Println(item.Name, item.Damage)

	/*championC := Models.Champion{Damage: 7, Name: "Pudge", HP: 26, Agility: 3, Item : item}
	db.Create(&championC)*/

	var items []Models.Item

	/*	db.Find(&items)
		db.Find(&heroes)*/
	heroes := make([]Models.Champion, 0)
	fmt.Println(heroes, " and ", items)

	heroes = append(heroes, championP, championI)
	PrepareToFight(heroes)
	fmt.Println("Agility ", heroes[0].GetAgility())
	Fight(heroes)
}

/*func main() {

	fmt.Println("Status .. Running...")

	axe := Models.Axe{Damage: 10}
	heroes := make([]Models.Hero, 0)
	pudge := Models.Pudge{
		Stats: Models.Stats{Damage: 6, Agility: 1, HP: 21}, Weapon: Models.Weapon(&axe),
	}
	invoker := Models.Invoker{
		Stats: Models.Stats{Damage: 7, Agility: 6, HP: 27}, Weapon: Models.Weapon(&axe),
	}


	heroes = append(heroes, &pudge, &invoker)

	PrepareToFight(heroes)
	Fight(heroes)

}*/

func PrepareToFight(heroes []Models.Champion) {

	for _, h := range heroes {
		//h.SetDamage(h.GetWeapon().GetDamage())
		if h.GetItem().GetName() == "Axe" {
			h.GetItem().GetUnique(&h)
		}
		//Not checking HP btw

	}

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Fight(heroes []Models.Champion) Models.Champion {

	fmt.Println("Fight begin ...")
	fmt.Println(heroes[0].GetName(), " VS ", heroes[1].GetName())
	var ChanseToMiss int

	for {
		ChanseToMiss = rand.Intn(10)

		time.Sleep(time.Second * 3)
		if heroes[0].GetHP() <= 0 {
			fmt.Println("Winner is ", heroes[1].GetName())
			return heroes[1]

		} else if heroes[1].GetHP() <= 0 {
			fmt.Println("Winner is ", heroes[0].GetName())
			return heroes[1]

		}

		if heroes[1].GetAgility() < ChanseToMiss {

			heroes[1].SetHP(-(heroes[0].GetDamage() + heroes[0].GetItem().GetDamage()))

			fmt.Println(heroes[0].GetName(), " Attacked ", heroes[1].GetName(), " with ", heroes[0].GetDamage()+heroes[0].GetItem().GetDamage(), " damage")

			fmt.Println(heroes[1].GetName(), " now have ", heroes[1].GetHP(), " HP")
			time.Sleep(time.Second)
			if heroes[1].GetHP() <= 0 {
				fmt.Println("Winner is ", heroes[0].GetName())
				return heroes[1]

			}

		} else {
			fmt.Println(heroes[0].GetName(), " just Missed with Roll", ChanseToMiss)
		}
		ChanseToMiss = rand.Intn(10)

		if heroes[0].GetAgility() < ChanseToMiss {
			heroes[0].SetHP(-heroes[1].GetDamage() - heroes[1].GetItem().GetDamage())

			fmt.Println(heroes[1].GetName(), " Attacked ", heroes[0].GetName(), " with ", heroes[1].GetDamage()+heroes[1].GetItem().GetDamage(), " damage")

			fmt.Println(heroes[0].GetName(), " now have ", heroes[0].GetHP(), " HP")

		} else {
			fmt.Println(heroes[1].GetName(), " just Missed with Roll", ChanseToMiss)
		}

		time.Sleep(time.Second)

	}

}
