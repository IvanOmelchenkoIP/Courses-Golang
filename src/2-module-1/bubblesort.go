package main

import (
	"fmt"
	"strconv"
)

func BubbleSort(numbers []int) {
	for ind, _ := range numbers {
		for i := 0; i < len(numbers) - ind - 1; i++ {
			if numbers[i] > numbers[i + 1] { Swap(numbers, i) }
		}
	}
}

func Swap(numbers []int, i int) {
	tmp := numbers[i]
	numbers[i] = numbers[i + 1]
	numbers[i + 1] = tmp
}

func PrintSlice(numbers []int) {
	for _, val := range numbers {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func main() {
	const maxNums = 10
	numbers := make([]int, 0, maxNums)
	
	var input string
	fmt.Println("Enter up to 10 integers") 
	for i := 0; i < maxNums; {
		fmt.Printf("Enter a number (X to stop): ")
		_, err := fmt.Scanln(&input)
		if err != nil { continue }

		if input == "X" || input == "x" { break }

		num, err := strconv.Atoi(input)
		if err != nil { continue }

		numbers = append(numbers, num)
		i++
	}

	if len(numbers) == 0 {
		fmt.Println("There was no data entered to read!")
		return
	}

	fmt.Printf("Unsorted slice: ")
	PrintSlice(numbers)

	BubbleSort(numbers)

	fmt.Printf("Sorted slice in order from the least to the greatest: ")
	PrintSlice(numbers)
}