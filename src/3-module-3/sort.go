package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"errors"
	"math"
	"sync"
	"sort"
)

func GetUserArray() ([]int, error) {
	fmt.Printf("Enter any number of integer numbers using space as a separator: ")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	if err != nil || len(input) == 0 {
		return make([]int, 0), errors.New("There was an error reading your input! Make sure your input corresponds to the prompt!")
	}
	inputArr := strings.Split(strings.TrimRight(string(input), " "), " ")

	numbers := make([]int, 0)
	for _, val := range inputArr {
		num, err := strconv.Atoi(val)
		if err != nil {
			return make([]int, 0), errors.New("Some of your input is not integers! Make sure your input corresponds to the prompt!")
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func PartitionArray(array []int, parts int) [][]int {
	partitionSlice := make([][]int, parts)

	arrLen := len(array)
	partPrecise := float64(float64(arrLen) / float64(parts))
	dec := partPrecise - float64(int(partPrecise))

	var lowCap int = 0
	var highCap int = 0
	for i := 0; i < parts; i++ {
		var partLen int
		switch {
		case dec < 0.5:
			if i < parts - 1 {
				partLen = int(math.Floor(partPrecise))
			} else {
				partLen = int(math.Ceil(partPrecise))
			}
		case dec == 0.5:
			if i < parts / 2 { 
				partLen = int(math.Ceil(partPrecise))
			} else {
				partLen = int(math.Floor(partPrecise))
			}
		default:
			if i < parts - 1 {
				partLen = int(math.Ceil(partPrecise))
			} else {
				partLen = int(math.Floor(partPrecise))
			}
		}
		highCap = lowCap + partLen
		if highCap > arrLen { 
			highCap = arrLen 
		}

		slice := array[lowCap:highCap]
		partitionSlice[i] = slice
		lowCap = highCap
	}
	return partitionSlice
}


func SortSubarray(array []int, i int, wg *sync.WaitGroup) {
	fmt.Printf("Subarray %d to be sorted: %d\n", i, array)
	sort.Ints(array)
	fmt.Printf("Subarray %d after sorting: %d\n", i, array)
	wg.Done()
}

func main() {
	const parts = 4

	numbers, err := GetUserArray()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Array of your entered integers:", numbers)

	if len(numbers) < parts {
		fmt.Println("Array could not be partitioned into 4 parts because its length is smaller that 4. It will be sorted without partitioning.")
		sort.Ints(numbers)
		fmt.Println("Sorted array:", numbers)
	} else {
		fmt.Println("Array will be partitioned into 4 parts.")
		partitionSlice := PartitionArray(numbers, parts)
		
		var wg sync.WaitGroup
		wg.Add(parts)
		for i, slice := range partitionSlice {
			go SortSubarray(slice, i, &wg)
		}

		wg.Wait()
		var i int
		for _, slice := range partitionSlice {
			for _, val := range slice {
				numbers[i] = val
				i++
			}
		}
		numbersSorted := make([]int, len(numbers))
		copy(numbersSorted, numbers)
		sort.Ints(numbersSorted)

		fmt.Println("The merged array of the sorted subarrays:", numbers)
		fmt.Println("The sorted merged array of the sorted subarrays:", numbersSorted)
	}
}