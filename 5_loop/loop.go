package main

import "fmt"

func main() {
	keepGoing := true
	answer := ""

	for keepGoing {
		fmt.Printf("Type command: ")
		fmt.Scan(&answer)

		if answer == "quit" {
			keepGoing = false
		}
	}

	fmt.Println("Program exit")
}
