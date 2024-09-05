package main

import (
	"AchivmentGame/Database"
	"AchivmentGame/Models"
	"AchivmentGame/Server"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//ОСТАЛОСЬ СДЕЛАТЬ

/*
Переделать чемпионов, сделать колличество убийств, сделать bool переменная killable
по дефолту false, если герой побеждает и не умирает то чекается кол-во убийств и проверяется killable
если killable false а сейчас уже 3 убийства, то присваивается Tripple Kill. Если его убили, то присваивается
killable = true, а при следующем бою, если боец победил и killable = true, то оно меняется на false.
*/
func main() {
	db, err := Database.InitDB()
	http.HandleFunc("/", Server.Handle)

	r := mux.NewRouter()
	Server.RegisterRoutes(r)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Models.Champion{})
	db.AutoMigrate(&Models.Item{})

	var championI []Models.Champion

	//Test adding champion
	/*championT := Models.Champion{ Damage: 12, Name: "Human", HP: 10, Agility: 7}
	db.Create(&championT)*/

	db.Order("RANDOM()").Limit(2).Find(&championI)
	for _, champion := range championI {

		fmt.Println("FOR ORDER", champion.GetName())

	}

	winner := Fight(championI)

	//add Kill stat to CHAMPION
	initWinner(db, &winner)

}

func postWinner(winner *Models.Champion) {
	var data = url.Values{}
	data.Add("name", winner.Name)
	data.Add("kills", strconv.Itoa(winner.GetKills()))
	data.Add("killable", strconv.FormatBool(winner.Killable))
	resp, err := http.PostForm("http://localhost:8090/winner", data)
	if err != nil {
		log.Fatalf("Failed to send winner: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Winner successfully sent")
	} else {
		log.Printf("Failed to send winner, status code: %d", resp.StatusCode)
	}
}

func initWinner(db *gorm.DB, winner *Models.Champion) {
	defer postWinner(winner)

	if winner.Killable == true {
		winner.Killable = false
	}

	db.Model(winner).Update("kills", winner.AddKill())
	db.Model(winner).Update("killable", winner.Killable)

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

		time.Sleep(time.Second * 1)
		if heroes[0].GetHP() <= 0 {
			fmt.Println("Winner is ", heroes[1].GetName())

			return heroes[1]

		} else if heroes[1].GetHP() <= 0 {
			fmt.Println("Winner is ", heroes[0].GetName())

			return heroes[0]

		}

		if heroes[1].GetAgility() < ChanseToMiss {

			heroes[1].SetHP(-heroes[0].GetDamage())

			fmt.Println(heroes[0].GetName(), " Attacked ", heroes[1].GetName(), " with ", heroes[0].GetDamage(), " damage")

			fmt.Println(heroes[1].GetName(), " now have ", heroes[1].GetHP(), " HP")
			time.Sleep(time.Second)
			if heroes[1].GetHP() <= 0 {
				fmt.Println("Winner is ", heroes[0].GetName())
				return heroes[0]

			}

		} else {
			fmt.Println(heroes[0].GetName(), " just Missed with Roll", ChanseToMiss)
		}
		ChanseToMiss = rand.Intn(10)

		if heroes[0].GetAgility() < ChanseToMiss {
			heroes[0].SetHP(-heroes[1].GetDamage())

			fmt.Println(heroes[1].GetName(), " Attacked ", heroes[0].GetName(), " with ", heroes[1].GetDamage(), " damage")

			fmt.Println(heroes[0].GetName(), " now have ", heroes[0].GetHP(), " HP")

		} else {
			fmt.Println(heroes[1].GetName(), " just Missed with Roll", ChanseToMiss)
		}

		time.Sleep(time.Second)

	}

}
