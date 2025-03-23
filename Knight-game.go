package main

import "fmt"

//TODO:

//!!!Upgrade inn and tavern inside town, room for 1, 7, 30 nights, improve dungeon for town
//!!!add dungeon in forest, add HP and DMG system to this game
//!!!Make Exit method from tavern for example
//!!Make the choices more difficult, add different rooms, food and drinks for tavern
//!!!Make fight method better, more interesting and work correctly
//!!!Don't forget about error handling

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
func (c *BaseCharacter) Fight() {
	fmt.Printf("%s fights\n", c.Name)
}

//Sleep method
func (c *BaseCharacter) Sleep() {
	fmt.Printf("%s sleep\n", c.Name)
}

//Drink method
func (c *BaseCharacter) Drink() {
	fmt.Printf("%s drinks\n", c.Name)
}

//Eat method
func (c *BaseCharacter) Eat() {
	fmt.Printf("%s eats food\n", c.Name)
}

//Fight overridden method for Warrior
func (k *Knight) Fight() {
	fmt.Printf("%s (Knight) hit the enemy with a stick\n", k.Name)
}

//Fight overridden method for Paladin
func (p *Paladin) Fight() {
	fmt.Printf("%s (Paladin) hit the enemy with a holy sword\n", p.Name)
}

//Fight overridden method for Mage
func (m *Mage) Fight() {
	fmt.Printf("%s (Mage) cast a fireball\n", m.Name)
}

func (r *Priest) Fight() {
	fmt.Printf("%s (Priest) use the holy hammer\n", r.Name)
}

//Choose location func
func (c *BaseCharacter) ChooseLocation() {
	var location string

	fmt.Print("Where do you want to go? (Forest/Town/Tavern)")
	_, err := fmt.Scan(&location)
	if err != nil {
		fmt.Println("input error", err)
		return
	}

	switch location {
	case "forest":
		c.GoToForest()
	case "town":
		c.GoToTown()
	case "tavern":
		c.GoToTavern()
	default:
		fmt.Print("Unknown location. Use fast travel to a familiar place")
		c.ChooseLocation()
	}
}

//Method for going to the forbidden forest
func (c *BaseCharacter) GoToForest() {
	fmt.Printf("%s goes to the Forbidden forest...\n", c.Name)
	fmt.Print("The monster has attacked you!\n")

	c.Fight()
}

//Method for going to the town
func (c *BaseCharacter) GoToTown() {
	fmt.Printf("%s goes to the Town...\n", c.Name)

	//Choose a place to visit in town
	var choosePlace string

	fmt.Print("Where do you want to go? (dungeon/inn/tavern)")
	fmt.Scan(&choosePlace)

	switch choosePlace {
	case "dungeon":
		fmt.Printf("%s goes to the dungeon...\n", c.Name)

	case "inn":
		fmt.Printf("%s goes to the Mermaid's inn...\n", c.Name)
	case "tavern":
		fmt.Printf("%s goes to the Brick's tavern...\n", c.Name)
	default:
		fmt.Print("Unknown location. Choose from familiar places\n")
	}

	//!!! Maybe delete this
	/*
	var actionInside string
	fmt.Scan(&actionInside)

	fmt.Print("What do you want to do? (Drink/Sleep)")

	switch actionInside {
	case "drink":
		c.Drink()
	case "sleep":
		c.Sleep()
	default:
		fmt.Print("Unknown action. Please try again\n")
	}*/
}

//Method for going to the Olaf's tavern
func (c *BaseCharacter) GoToTavern() {
	fmt.Printf("%s goes to the Olaf's tavern...\n", c.Name)

	var tavernAction string

	fmt.Print("Olaf: Hey, traveler! Do you want to drink a cold beer or a room for a night? (Eat/Drink/Sleep)")

	//Error handling
	_, err := fmt.Scan(&tavernAction)
	if err != nil {
		fmt.Println("input error:", err)
	}

	switch tavernAction {
	case "drink":
		fmt.Printf("Olaf: Here's your beer %s\n", c.Name)
		c.Drink()
	case "sleep":
		fmt.Print("Olaf: your room is second on the right side of second floor\n")
		c.Sleep()
	case "eat":
		fmt.Print("Waitress: Here's your food \n")
		c.Eat()
	default:
		fmt.Print("Olaf: We don't offer this option\n")
	}
}

func StartGame() {
	var name, characterClass string

	//Type a name of your character
	fmt.Print("Type name of your character: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("input error:", err)
		return
	}

	//Choose a class of character
	fmt.Print("Choose class of your character: (knight/paladin/mage/priest)")
	_, err = fmt.Scan(&characterClass)
	if err != nil {
		fmt.Println("input error:", err)
		return
	}

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
		character = &Knight{BaseCharacter{Name: name}}
	}
	character.ChooseLocation()
}

func main() {
	StartGame()
}
