package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"errors"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
	food string
	locomotion string
	noise string
}

func (cow Cow) Eat() {
	fmt.Printf("Food eaten by %s the cow is %s.\n\n", cow.name, cow.food)
} 

func (cow Cow) Move() {
	fmt.Printf("Locomotion method of %s the cow is %s.\n\n", cow.name, cow.locomotion)
} 

func (cow Cow) Speak() {
	fmt.Printf("Spoken sound of %s the cow is %s.\n\n", cow.name, cow.noise)
}

type Bird struct {
	name string
	food string
	locomotion string
	noise string
}

func (bird Bird) Eat() {
	fmt.Printf("Food eaten by %s the bird is %s.\n\n", bird.name, bird.food)
} 

func (bird Bird) Move() {
	fmt.Printf("Locomotion method of %s the bird is %s.\n\n", bird.name, bird.locomotion)
} 

func (bird Bird) Speak() {
	fmt.Printf("Spoken sound of %s the bird is %s.\n\n", bird.name, bird.noise)
}

type Snake struct {
	name string
	food string
	locomotion string
	noise string
}

func (snake Snake) Eat() {
	fmt.Printf("Food eaten by %s the snake is %s.\n\n", snake.name, snake.food)
} 

func (snake Snake) Move() {
	fmt.Printf("Locomotion method of %s the snake is %s.\n\n", snake.name, snake.locomotion)
} 

func (snake Snake) Speak() {
	fmt.Printf("Spoken sound of %s the snake is %s.\n\n", snake.name, snake.noise)
}

func CreateNewAnimal(reqName, reqAnimal string, animalList map[string]Animal) error {
	switch reqAnimal {
	case "cow":
		animalList[reqName] = Cow{name: reqName, food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		animalList[reqName] = Bird{name: reqName, food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		animalList[reqName] = Snake{name: reqName, food: "mice", locomotion: "slither", noise: "hsss"}
	default:
		return errors.New("Can not create requested animal!")
	}
	return nil
}

func QueryAnimal(reqName, reqAction string, animalList map[string]Animal) error {
	var animal Animal = animalList[reqName]
	if animal == nil {
		return errors.New("There is no animal with the requested name!")
	}

	switch reqAction {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		return errors.New("There is no data on requested property of the animal!")
	}
	return nil
}

func main() {
	var animalList = make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a one of the commands to continue:")
		fmt.Println("'newanimal name type(cow, bird or snake)' to create a new animal")
		fmt.Println("'query name action(eat, move or speak)' to get information about the animal")
		fmt.Println("'X' to exit")

		fmt.Printf("> ")
		request, _, err := reader.ReadLine()
		if err != nil { 
			fmt.Println("There was an error reading your inut! Make sure your input corresponds to the prompt!")
			fmt.Println()
		}
		if strings.ToLower(string(request)) == "x" { break }

		queries := strings.Split(string(request), " ")
		if len(queries) != 3 {
			fmt.Println("Make sure there is correct amount of arguments in your request!")
			fmt.Println()
			continue
		}

		reqCommand := queries[0]
		switch reqCommand {
		case "newanimal":
			reqName := queries[1]
			reqAnimal := queries[2]
			err := CreateNewAnimal(reqName, reqAnimal, animalList)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
			} else {
				fmt.Printf("%s the %s was successfully created!\n\n", reqName, reqAnimal)
			}
		case "query":
			reqName := queries[1]
			reqAction := queries[2]
			err := QueryAnimal(reqName, reqAction, animalList)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
			}
		default:
			fmt.Println("There is no command that you requested!")
			fmt.Println()
		}
	}
}