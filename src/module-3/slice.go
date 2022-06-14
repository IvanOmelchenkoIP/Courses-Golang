package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	// In the task we are wanted to create an empty integer slice of size (length) 3
	// Creating an empty slice of size 3 would probably mean a slice with length 0 and capacity 3
	// However, in the task it is stated that the length must be 3
	// So it may also mean not an empty slice, but a slice of length 3 with 3 zero elements
	// I have decided to implement both slices
	
	var input string
	const initLen = 3
	sliceLen := make([]int, initLen) // slice with 3 zero elements with length 3
	sliceCap := make([]int, 0, initLen) // empty slice with underlying array length 3
	for i := 1;; {		
		fmt.Printf("Enter the integer to add to the slice (type X to exit): ")
		_, err := fmt.Scanln(&input)
		if err != nil { continue }
		if (input == "x" || input == "X") { break } 

		num, err := strconv.Atoi(input)
		if err != nil { continue }

		if i > initLen {
			sliceLen = append(sliceLen, num)
		} else {
			for ind, val := range sliceLen {
				if val == 0 {
					sliceLen[ind] = num
					break
				}
			}
		}
		sort.Ints(sliceLen)
		fmt.Println("Sorted slice with initial length 3 and zero values elements:", sliceLen) 

		sliceCap = append(sliceCap, num)
		sort.Ints(sliceCap)
		fmt.Println("Sorted initially empty slice with initial capacity 3:", sliceCap) 
		
		i++
	}
}