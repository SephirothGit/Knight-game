package main

import (
	"fmt"
	"strings"
)

//TODO:

//!!!Fix loop in Mermaid's inn
//!!!Make more fights
//!!!Make fight method more interesting
//!!!Add different rooms, food and drinks for tavern, action to the dungeon, inn, tavern

// Character interface that defines methods for all characters
type Character interface {
	Fight()
	Sleep()
	Drink()
	Eat()
	ChooseLocation()
}

type BaseCharacter struct {
	Name   string
	HP     int
	MaxHP  int
	DMG    int
	Gold   int
	Silver int
}

type Knight struct {
	BaseCharacter
}

// Structure for Paladin
type Paladin struct {
	BaseCharacter
}

// Structure for Mage
type Mage struct {
	BaseCharacter
}

// Structure for Priest
type Priest struct {
	BaseCharacter
}

// Fight method
func (c *BaseCharacter) Fight() {
	fmt.Printf("%s fights\n", c.Name)
}

// Sleep and restore full HP method
func (c *BaseCharacter) Sleep() {
	c.HP = c.MaxHP
	fmt.Printf("%s sleep and restores full HP %d/%d\n", c.Name, c.HP, c.MaxHP)
}

// Drink method
func (c *BaseCharacter) Drink() {
	if c.Silver >= 1 {
		c.Silver -= 1
		fmt.Printf("%s drinks a beer (-1 silver coin)\n", c.Name)
	} else {
		fmt.Println("Not enough silver!")
	}
}

// Eat method
func (c *BaseCharacter) Eat() {
	if c.Silver >= 2 {
		c.Silver -= 2
		c.HP += 50

		if c.HP > c.MaxHP {
			c.HP = c.MaxHP
		}
		fmt.Printf("%s eats food and restores 50 HP (-2 silver coin)\n", c.Name)
	} else {
		fmt.Println("Not enough silver")
	}
}

// Fight overridden method for Warrior
func (k *Knight) Fight() {
	fmt.Printf("%s (Knight) hit the enemy with a stick for %d DMG\n", k.Name, k.DMG)
}

// Fight overridden method for Paladin
func (p *Paladin) Fight() {
	fmt.Printf("%s (Paladin) hit the enemy with a holy sword for %d DMG\n", p.Name, p.DMG)
}

// Fight overridden method for Mage
func (m *Mage) Fight() {
	fmt.Printf("%s (Mage) cast a fireball for %d DMG\n", m.Name, m.DMG)
}

func (r *Priest) Fight() {
	fmt.Printf("%s (Priest) use the holy hammer for %d DMG\n", r.Name, r.DMG)
}

// Choose location func
func (c *BaseCharacter) ChooseLocation() {
	for {
		var location string

		fmt.Print("Where do you want to go? (Forest/Town/Tavern/Exit)")
		_, err := fmt.Scan(&location)
		if err != nil {
			fmt.Println("input error", err)
			return
		}

		switch strings.ToLower(location) {
		case "forest":
			c.GoToForest()
			return
		case "town":
			c.GoToTown()
			return
		case "tavern":
			c.GoToTavern()
			return
		case "exit":
			fmt.Print("Goodbye")
			return
		default:
			fmt.Print("Unknown location. Use fast travel to a familiar place\n")
			c.ChooseLocation()
		}
	}
}

// Method for going to the forbidden forest
func (c *BaseCharacter) GoToForest() {
	fmt.Printf("%s goes to the Forbidden forest...\n", c.Name)

	//Improved fight system
	enemyHP := 50
	fmt.Printf("A monster (HP: %d) has attacked you!\n", enemyHP)

	for {
		fmt.Printf("%s HP: %d/%d | Monster HP %d\n", c.Name, c.HP, c.MaxHP, enemyHP)
		fmt.Print("What will you do? (Fight/Run)")

		var action string
		fmt.Scan(&action)

		switch strings.ToLower(action) {
		case "fight":
			fmt.Printf("%s atacks the monster!\n", c.Name)
			enemyHP -= c.DMG

			if enemyHP <= 0 {
				fmt.Println("You defeated the monster!\n")
				return
			}
			//Monster fights
			monsterDMG := 10
			c.MaxHP -= monsterDMG
			fmt.Printf("Monster hits you for %d damage!\n", monsterDMG)

			if c.HP <= 0 {
				fmt.Println("You died")
				return
			}
		case "run":
			fmt.Print("You run away from the monster!")
			return
		default:
			fmt.Print("Unknown command")
		}
	}
}

// Method for going to the town
func (c *BaseCharacter) GoToTown() {
	fmt.Printf("%s goes to the Town...\n", c.Name)

	for {
		//Choose a place to visit in town
		var choosePlace string

		fmt.Print("Which place you want to visit? (Dungeon/Inn/Tavern/Exit)")
		_, err := fmt.Scan(&choosePlace)
		if err != nil {
			fmt.Println("input error:", err)
		}
		switch strings.ToLower(choosePlace) {
		case "dungeon":
			c.GoToDungeon()
			return
		case "inn":
			fmt.Printf("%s goes to the Mermaid's inn...\n", c.Name)
			c.GoToInn()
			return
		case "tavern":
			fmt.Printf("%s goes to the Brick's tavern...\n", c.Name)
			c.GoToTavernB()
			return
		case "exit":
			fmt.Print("You came out\n")
			c.ChooseLocation()
			return
		default:
			fmt.Print("Unknown location. Choose from familiar places\n")
		}
	}
}

// Method for going to the Olaf's tavern
func (c *BaseCharacter) GoToTavern() {
	fmt.Printf("%s goes to the Olaf's tavern...\n", c.Name)

	//Loop for actions
	for {

		var tavernAction string

		fmt.Print("Olaf: Hey, traveler! Do you want to drink a cold beer or a room for a night? (Eat/Drink/Sleep/Exit)")

		//Error handling
		_, err := fmt.Scan(&tavernAction)
		if err != nil {
			fmt.Println("input error:", err)
		}

		switch strings.ToLower(tavernAction) {
		case "drink":
			fmt.Printf("Olaf: Here's your beer %s\n", c.Name)
			c.Drink()
		case "sleep":
			fmt.Print("Olaf: your room is second on the right side of second floor\n")
			c.Sleep()
		case "eat":
			fmt.Print("Waitress: Here's your food \n")
			c.Eat()
		case "exit":
			fmt.Print("Olaf: Goodbye, traveler!\n")
			c.ChooseLocation()
			return
		default:
			fmt.Print("Olaf: We don't offer this option\n")
		}
	}
}

func (c *BaseCharacter) GoToTavernB() {
	fmt.Printf("%s goes to the Brick's Tavern...\n", c.Name)

	for {
		var tavernActionB string

		fmt.Print("Brick: Good day, traveler! do you want to drink a glass of cold beer with a snacks? (Drink/Eat/Exit)")
		_, err := fmt.Scan(&tavernActionB)
		if err != nil {
			fmt.Println("input err", err)
		}

		switch strings.ToLower(tavernActionB) {
		case "drink":
			fmt.Printf("Brick: here's your beer %s\n", c.Name)
			c.Drink()
		case "eat":
			fmt.Printf("Brick: here's your snacks %s\n", c.Name)
			c.Eat()
		case "exit":
			fmt.Print("Brick: come back anytime!\n")
		default:
			fmt.Print("UNknown action. Choose from (Drink/Eat/Exit)\n")
		}
	}
}

func (c *BaseCharacter) GoToDungeon() {

	for {
		var dungeonAction string

		fmt.Printf("%s went down into the dungeon...\n", c.Name)

		fmt.Print("Choose floor: (First/Second/Third/Teleport)")

		_, err := fmt.Scan(&dungeonAction)
		for err != nil {
			fmt.Println("input error:", err)
		}

		switch strings.ToLower(dungeonAction) {
		case "first":
			fmt.Printf("%s went down to the 1 floor of the Ancient Elven Dungeon\n", c.Name)
		case "second":
			fmt.Printf("%s went to the middle floor of the Ancient Elven Dungeon\n", c.Name)
		case "third":
			fmt.Printf("%s went to the lower floor of the Ancient Elven Dungeon\n", c.Name)
		case "teleport":
			fmt.Print("Choose a place to teleport:\n")
			c.ChooseLocation()
			return
		default:
			fmt.Print("Unknown command. Choose from known floors\n")
		}
	}
}
func (c *BaseCharacter) GoToInn() {
	fmt.Print("Welcome to the Mermaid's Inn!\n")

	for {
	var innDuration string
	var innAction string
	fmt.Print("Freya: We offer best rooms in Fawn! For how long you want to stay in town? (1/7/30)")

	_, err := fmt.Scan(&innDuration)
	if err != nil {
		fmt.Println("input error:", err)
		continue
	}

	//Choose for how long you want to stay in Fawn
	switch strings.ToLower(innDuration) {
	case "1":
		fmt.Print("Freya: It will cost you 1 silver coin\n")

		if c.Silver >= 1 {
			c.Silver -= 1
			fmt.Print("1 silver coin was given away\n")
			c.Sleep()
		} else {
			fmt.Println("Not enough silver")
		}

	case "7":
		fmt.Print("Freya: It will cost you 7 silver coins\n")

		if c.Silver >= 7 {
			c.Silver -= 7
			fmt.Print("7 silver coins were given away\n")
			c.Sleep()
		} else {
			fmt.Println("Not enough silver")
		}
	case "30":
		fmt.Print("Freya: Great choise! It will cost you only 27 silver coins!\n")

		if c.Silver >= 27 {
			c.Silver -= 27
			fmt.Print("27 silver coins were given away\n")
			c.Sleep()
		} else {
			fmt.Println("Not enough silver!")
		}

		for {
			fmt.Printf("Freya: %s, do you want to eat and drink, or maybe you want to hear rumors? (Drink/Eat/Rumors/Exit)", c.Name)

			_, err = fmt.Scan(&innAction)
			if err != nil {
				fmt.Println("input error:", err)
			}

			switch strings.ToLower(innAction) {
			case "drink":
				c.Drink()
			case "eat":
				c.Eat()
			case "rumors":
				fmt.Print("Freya: You want to hear the rumors, ok it will cost 2 silver coins\n")

				if c.Silver >= 2 {
					c.Silver -= 2
					fmt.Print("2 silver coins were given away\n")
				}
				fmt.Print("Freya: Some strangers that were here two days ago asked travelers about man with a sword named Steel Rose. I heard that he stealed this item from the head of a royal guard of Shangri-La. A reward of 20 gold coins has been offered for the return of this sword. If you fast and smart enough to overtake those men, you can get this reward.\n")
			case "exit":
				fmt.Print("Freya: come back whenever you want to stay in town, traveler!\n")
				c.ChooseLocation()
				return
			}
		}
	}
}

func StartGame() {
	var name, characterClass string

	//Type a name of your character
	fmt.Print("Give a name for your character: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("input error:", err)
		return
	}

	//Choose a class of character
	fmt.Print("Choose class of your character: (Knight/Paladin/Mage/Priest)")
	_, err = fmt.Scan(&characterClass)
	if err != nil {
		fmt.Println("input error:", err)
		return
	}

	var character Character

	switch strings.ToLower(characterClass) {
	case "knight":
		character = &Knight{BaseCharacter{Name: name, HP: 180, MaxHP: 180, DMG: 20, Gold: 10, Silver: 120}}
	case "paladin":
		character = &Paladin{BaseCharacter{Name: name, HP: 220, MaxHP: 220, DMG: 24, Gold: 13, Silver: 146}}
	case "mage":
		character = &Mage{BaseCharacter{Name: name, HP: 120, MaxHP: 120, DMG: 36, Gold: 34, Silver: 293}}
	case "priest":
		character = &Priest{BaseCharacter{Name: name, HP: 80, MaxHP: 80, DMG: 8, Gold: 2, Silver: 91}}
	default:
		fmt.Println("Unknown class. Create Knight by default")
		character = &Knight{BaseCharacter{Name: name, HP: 180, MaxHP: 180, DMG: 20, Gold: 10, Silver: 120}}
	}
	character.ChooseLocation()
}

func main() {
	StartGame()
}
