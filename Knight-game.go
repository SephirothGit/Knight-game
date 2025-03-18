package main

import "fmt"

//Character interface that defines methods for all characters
type Knight interface {
	Fight()
	Sleep()
	Drink()
	Eat()
	ChooseLocation()
}

type BaseCharacter struct {
	Name string // !!!Add a name insert!!!
}

//Structure for Paladin
type Paladin struct {
	BaseCharacter
} 
//Structure for Mage
type Mage struct {
	BaseCharacter
} 

//Fight method
func (c BaseCharacter) Fight() { //!!!Without asterisk!!!
	fmt.Printf("%s fights\n", c.Name)
}

//Sleep method
func (c *BaseCharacter) Sleep() {
	fmt.Printf("%s sleep\n", c.Name)
}

//Drink method
func (c *BaseCharacter) Drink() {
	fmt.Printf("%s drink\n", c.Name)
}

//Eat method
func (c BaseCharacter) Eat() {
	fmt.Printf("%s eat\n", c.Name)
}

//Fight overridden method for Paladin
func (p *Paladin) Fight() {
	fmt.Printf("%s (Paladin) hit the enemy with a holy sword", p.Name)
} 

//Fight overridden method for Mage
func (m *Mage) Fight() {
	fmt.Printf("%s (Mage) cast a fireball", m.Name)
}

//Choose location func
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
		c.GoToTavern()
	default:
		fmt.Println("Unknown location. Use fast travel to a familiar place")
		c.ChooseLocation()
	}
}


//Method for going to the forest
func(c BaseCharacter) GoToForest() {
	fmt.Printf("%s goes to the Forbidden forest...\n", c.Name)
	fmt.Println("The monster attacked you!")

	c.Fight() //!!!Make fight method better and more interesting
}

//Method for going to the town
func(c BaseCharacter) GoToTown() {
	fmt.Printf("%s goes to the Town...\n", c.Name) //!!!Create inn and tavern inside town, room for 1, 7, 30 nights

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

//Method for going to the tavern
func(c BaseCharacter) GoToTavern() {
	fmt.Printf("%s goes to the Olaf's tavern...\n", c.Name) //!!Make the choices more difficult, add different rooms, food and drinks

	var tavernAction string
	fmt.Scan(&tavernAction)

	fmt.Print("Olaf: Hey, traveler! Do you want to drink a cold beer or a room for a night? (Eat/Drink/Sleep)")

	switch tavernAction {
	case "drink":
		fmt.Printf("Olaf: Here's your beer %s", c.Name)
		c.Drink()
	case "sleep":
		fmt.Print("Olaf: your room is second on the right side of second floor")
		c.Sleep()
	case "eat":
		fmt.Print("Waitress: Here's your food")
		c.Eat()
	default:
		fmt.Println("Olaf: We don't offer this option")
	}

}

func main() {

}
