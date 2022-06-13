package main

import (
	"fmt"
	"strings"
)

var str string

func main() {
	fmt.Printf("Enter a string: ")

	var temp string
	for flagEnd := false; flagEnd != true; {
		num, err := fmt.Scanf("%s", &temp)
		if num < 1 || err != nil { 
			flagEnd = true
			str = strings.TrimSuffix(str, " ") 
		} else {
			str += temp + " "
		}
	}

	str = strings.ToLower(str)
	var iInd int = strings.Index(str, "i")
	var nInd int = strings.LastIndex(str, "n")
	var aInd int = strings.Index(str, "a")
	if iInd == 0 && nInd == len(str) - 1 && aInd != -1{
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}