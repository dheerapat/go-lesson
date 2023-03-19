package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func main() {
	no1, _ := strconv.Atoi(os.Args[1]) // Atoi = ASCII to Integer
	no2, _ := strconv.Atoi(os.Args[2]) // os.Args = command-line argument at runtime

	sum := add(no1, no2)
	fmt.Println(sum)
}
