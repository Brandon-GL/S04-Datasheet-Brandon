package player

import (
	"fmt"
	"buffio"
	"io"
	"log"
	"os"
)

type Player struct {
	Nom string
	Pseudo string
	Age int
	Health int
	Mana int
}

var {
	players = make(map[string]*Player)
}

func player() {

}

func addPlayer() {
	fmt.Printf("\n Saisissez votre nom : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	nom := scan.Text()
	log.Printf("Nom : %v \n", nom)

	fmt.Printf("\n Saisissez votre pseudo : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	pseudo := scan.Text()
	log.Printf("Pseudo : %v \n", pseudo)

	fmt.Printf("\n Saisissez votre age : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	age := scan.Text()
	log.Printf("Age : %v \n", age)

	t := Player{
		Nom: nom,
		Pseudo: pseudo,
		Age: age,
		Health: 100,
		Mana: 100,
	}

	players[nom] = &t
}

func (t Player) Display() {
	fmt.Printf("\n Nom: %v, \n Pseudo: %v, \n Age: %v, \n Health: %v, \n Mana: %v \n", t.Nom, t.Pseudo, t.Age, t.Health, t.Mana)
}

func save() {
	os.WriteFile("%v.yml", nom, t.Player, 0600)
}

func del() {
	
}