package utils

import (
	"fmt"
)

func InputInteger(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scanln(&input)

	return input
}
