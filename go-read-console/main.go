package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go read console input sample app")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()

		if len(text) != 0 {
			fmt.Println("You entered:", text)
		} else {
			break
		}

	} //for

	fmt.Println("Finished")
} //main
