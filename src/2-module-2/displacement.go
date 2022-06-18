package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(acceleration, initVelocity, initDisplacement float64) func (float64) float64 {
	CountDisplacement := func (time float64) float64 {
		displacement := 0.5 * acceleration * time * time + initVelocity * time + initDisplacement
		return displacement
	}
	return CountDisplacement
}

func GetInitInput (message string) float64 {
	for {
		fmt.Printf("%s ", message)
		var input string
		fmt.Scanln(&input)

		num, err := strconv.ParseFloat(input, 64)
		if err != nil { continue }
		return num
	}
}

func main() {	
	fmt.Println("The program counts displacement by formula: 1/2 * acceleration * time^2 + initial_velocity * time + initial_displacement")
	fmt.Println("Enter initial data (floating point numbers):")
	acceleration := GetInitInput("Enter the acceleration:")
	initVelocity := GetInitInput("Enter the initial velocity:")
	initDisplacement := GetInitInput("Enter the initial displacement:")

	fn := GenDisplaceFn(acceleration, initVelocity, initDisplacement)
	fmt.Println("Initial data has been recorded")

	for {
		var input string
		fmt.Printf("Enter time to count displacement using initial data (enter X to exit the loop): ")
		fmt.Scanln(&input)

		if input == "X" || input == "x" { break }
		time, err := strconv.ParseFloat(input, 64)
		if err != nil { continue }

		displacement := fn(time)
		fmt.Println("Displacement travelled after time t:", displacement)
	}
}