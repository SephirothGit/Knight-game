package main

import "fmt"

//!!!add dungeon in forest, HP and DMG

//Character interface that defines methods for all characters
type Character interface {
	Fight()
	Sleep()
	Drink()
	Eat()
	ChooseLocation()
}

type BaseCharacter struct {
	Name string
}

type Knight struct {
	BaseCharacter
}

//Structure for Paladin
type Paladin struct {
	BaseCharacter
}

//Structure for Mage
type Mage struct {
	BaseCharacter
}

//Structure for Priest
type Priest struct {
	BaseCharacter
}

//Fight method
func (c *BaseCharacter) Fight() error {  //!!!Error handling!!!
	fmt.Printf("%s fights\n", c.Name)
	return nil
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

//Fight overridden method for Warrior
func (k *Knight) Fight() {
	fmt.Printf("%s (Knight) hit the enemy with a stick", k.Name)
}

//Fight overridden method for Paladin
func (p *Paladin) Fight() {
	fmt.Printf("%s (Paladin) hit the enemy with a holy sword", p.Name)
}

//Fight overridden method for Mage
func (m *Mage) Fight() {
	fmt.Printf("%s (Mage) cast a fireball", m.Name)
}

func (r *Priest) Fight() {
	fmt.Printf("%s (Priest) use the holy hammer", r.Name)
}

//Choose location func
func (c *BaseCharacter) ChooseLocation() {
	var location string
	
	fmt.Print("Where do you want to go? (Forest/Town/Tavern)")          //!!!Fix error handling
	//Error handling
//	_, err := fmt.Scanln(&location)
//	if err != nil {
//		return fmt.Errorf("input error: %v", err)
	}

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
		fmt.Print("Unknown location. Use fast travel to a familiar place")
		c.ChooseLocation()
	}
}

//Method for going to the forest
func (c BaseCharacter) GoToForest() {
	fmt.Printf("%s goes to the Forbidden forest...\n", c.Name)
	fmt.Println("The monster attacked you!")

	c.Fight() //!!!Make fight method better and more interesting
}

//Method for going to the town
func (c BaseCharacter) GoToTown() {
	fmt.Printf("%s goes to the Town...\n", c.Name) //!!!Upgrade inn and tavern inside town, room for 1, 7, 30 nights, improve dungeon

	//Choose a place to visit in town
	var choosePlace string
	fmt.Scan(&choosePlace)

	fmt.Print("Where do you want to go? (dungeon/inn/tavern)")

	switch choosePlace {
	case "dungeon":
		fmt.Printf("%s goes to the dungeon...", c.Name)
	case "inn":
		fmt.Println("%s goes to the Mermaid's inn...", c.Name)
	case "tavern":
		fmt.Println("%s goes to the Brick's tavern...", c.Name)
	default:
		fmt.Print("Unknown location. Choose from familiar places")
	}

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
func (c BaseCharacter) GoToTavern() {
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

func StartGame() {
	var name, characterClass string

	//Type a name of your character
	fmt.Print("Type name of your character: ")
	fmt.Scanln(&name)
	//Choose a class of character
	fmt.Print("Choose class of your character: (knight/paladin/mage/priest)")
	fmt.Scan(&characterClass)

	var character Character

	switch characterClass {
	case "knight":
		character = &Knight{BaseCharacter{Name: name}}
	case "paladin":
		character = &Paladin{BaseCharacter{Name: name}}
	case "mage":
		character = &Mage{BaseCharacter{Name: name}}
	case "priest":
		character = &Priest{BaseCharacter{Name: name}}
	default:
		fmt.Println("Unknown class. Create knight by default")
	}
	character.ChooseLocation()
}

func main() {
	StartGame()
}
