package main

import (
	"fmt"
	"encoding/json"
)

var name string
var addr string

func main() {
	fmt.Printf("Enter the name: ")
	fmt.Scanln(&name)
	fmt.Printf("Enter the address: ")
	fmt.Scanln(&addr)

	usrInfo := make(map[string]string)
	usrInfo["name"] = name 
	usrInfo["address"] = addr

	usrInfoJSON, err := json.Marshal(usrInfo)
	if err == nil {
		fmt.Println("JSON object, created from input data:", string(usrInfoJSON))
	}
}
