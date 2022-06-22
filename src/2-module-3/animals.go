package main

import (
	"fmt"	
	"strings"
	"errors"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal Animal) Eat() {
	fmt.Println("Food eaten by the animal:", animal.food)
}

func (animal Animal) Move() {
	fmt.Println("Locomotion method of the animal:", animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println("Spoken sound of the animal:", animal.noise)
}

func GetAnimalType(reqAnimal string) (Animal, error) {
	var animal Animal
	var err error
	switch reqAnimal {
	case "cow": 
		animal = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird": 
		animal = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	case "snake": 
		animal = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	default:
		err = errors.New("There is no data on requested animal!")
	}
	return animal, err
}

func main() {
	for {
		fmt.Println("Enter animal type (cow, bird, snake) and the action (eat, move, speak) as a request or X to exit")
		fmt.Printf("> ")
		var reqPrimary string
		var reqAction string
		fmt.Scan(&reqPrimary, &reqAction)
		if strings.ToLower(reqPrimary) == "x" { break }

		animal, err := GetAnimalType(reqPrimary)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch reqAction {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("There is no data on requested property of the animal!")
			continue
		}
	}
}