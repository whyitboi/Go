//Decimal to Binary

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userInput int64
	fmt.Print("Please enter an integer: ")
	fmt.Scan(&userInput)

	result := strconv.FormatInt(userInput, 2)
	fmt.Print(userInput, " in Binary is: ", result)
}
