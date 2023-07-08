package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Go json sample app")

	doc1 := DocumentInfo{Name: "Andrew", Id: 10, Description: []string{"Hello", "World!"}}

	doc1.Description = []string{"apple", "pie", "Hello"}

	jsonBytes, _ := json.Marshal(doc1)
	jsonText := string(jsonBytes)

	fmt.Println(jsonText)

	fmt.Println("Finished")
} //main
