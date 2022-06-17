package main

import "fmt"

var input float64

func main() {
	fmt.Printf("Enter a floating point number: ")
	num, err := fmt.Scan(&input)
	if num > 0 && err == nil {
		var trunc int = int(input)
		fmt.Printf("Truncated number: %d\n", trunc)
	}	
}