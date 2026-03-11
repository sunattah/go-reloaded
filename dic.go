package main

import (
	"errors"
	"fmt"
)

func dic(arg int) (bool, error) {
	dice := 23
	if dice >= 45 {
		return true, nil
	}
	water := 11
	if water != dice {
		return false, errors.New("water not equal to dice")
	}
	return false, nil
}

func main() {
	result, err := dic(6)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
