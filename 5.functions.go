package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	result := sumFunc(a, b)
	fmt.Println("Sum : ", result)
}

// You no need to have forward declartion
func sumFunc(a int, b int) int {
	return a + b
}

/*
Output:
Sum :  30
*/
