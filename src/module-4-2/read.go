package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

const fieldLen = 20

type Name struct {
	fname string
	lname string
} 

var filename string

func main() {
	fmt.Printf("Enter a file name: ")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File with that name could not be found!")
		return
	}

	fmt.Println("File successfully has been opened. Reading the contents:")
	scanner:= bufio.NewScanner(file)
	uSlice := make([]Name, 0)
	for scanner.Scan() {
		nameStr := scanner.Text()
		if len(nameStr) == 0 { break } 

		nameArr := strings.Split(string(nameStr), " ")
		first := nameArr[0]
		last := nameArr[1]
		
		// Since in the task it is said that each field will be a string of size 20 (characters),
		// I am assuming that the longer strings should be truncated to meet the size of 20 characters
		if len(first) > fieldLen {
			first = first[0:fieldLen]
		}
		if len(last) > fieldLen {
			last = last[0:fieldLen]
		}

		uName := Name{fname: first, lname: last}
		uSlice = append(uSlice, uName)
	}
	file.Close()

	if (len(uSlice) == 0) {
		fmt.Println("Contents of the file could not be read!")
		return
	}

	for _, name := range uSlice {
		fmt.Println(name.fname, name.lname)
	}
}