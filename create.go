package main

import (
	"fmt"
	"os"
)

func main() {
	file, num := os.Create("sample.txt")

	if num != nil {
		fmt.Println("Error: ", num)
	}
	fmt.Print(file)
}
