package main

import "fmt"

type Knight interface {
	Fight()
	Sleep()
	Drink()
	ChooseLocation()
}

type BaseCharacter struct {
	Name string // !!!Add a name insert!!!
}

func (c BaseCharacter) Fight() { //!!!Without asterisk!!!
	fmt.Printf("%s fights\n", c.Name)
}

func (c *BaseCharacter) Sleep() {
	fmt.Printf("%s sleep\n", c.Name)
}

func (c *BaseCharacter) Drink() {
	fmt.Printf("%s drink\n", c.Name)
}

func (c *BaseCharacter) ChooseLocation() {
	var location string

	fmt.Print("Where you want to go? (Forest/Town/Tavern)")
	fmt.Scanln(&location)

	switch location {
	//case "Forbidden woods":
	//fmt.Println("You are going to Forbidden woods")
	case "forest":
		c.GoToForest()
	case "town":
		//fmt.Println("You are going to Town")
		c.GoToTown()
	//case "Olaf's Tavern":
	//	fmt.Println("You are going to Olaf's tavern")
	case "tavern":
	default:
		fmt.Println("Unknown location. Use fast travel to a familiar place")
		c.ChooseLocation()
	}
}

func(c BaseCharacter) GoToForest() {
	fmt.Printf("%s goes to the Forest...\n", c.Name)
	fmt.Println("The monster attacked you!")

	c.Fight()
}

func(c BaseCharacter) GoToTown() {
	fmt.Printf("%s goes to the Town...\n", c.Name)

	var action string
	fmt.Scan(&action)

	fmt.Print("What do you want to do? (Drink/Sleep)")

	switch action {
	case "drink":
		c.Drink()
	case "sleep":
		c.Sleep()
	default:
		fmt.Println("Unknown action. Please try again")
	}
}

func main() {

}
