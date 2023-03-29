package main

import "fmt"

func main() {
	fmt.Println("Enter player names (separated by space or newline):")
	var player1, player2 string
	fmt.Scan(&player1, &player2)
	fmt.Println("Player are: ", player1, player2)
}
