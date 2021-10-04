package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// We are going to use the random numbers, So we need to provide the seed value to random generator.
	// Otherwise it will generate the same number for each execution
	// We also need to provide the unique value, So we are using the 'time' as the seed for random generator.
	// This is very similar to other languages random generators like C++, Python
	rand.Seed(time.Now().UnixNano())

	// rand.Intn -> Generates the number upto given number.
	var roger, novak int = rand.Intn(100), rand.Intn(100)

	fmt.Printf("Roger Age %d \n", roger)
	fmt.Printf("Novak Age %d \n", novak)

	// The GoLang 'if..else' Block
	// Very similar to C if..else but no need of brackets around conditions
	// The curly brackes are mandatory, Even for single line inside the if block. Unlike C-Lang
	// The else must be in the below format. '} else {' - So careful :)
	if roger > novak {
		fmt.Println("Roger is elder")
	} else {
		if roger < novak {
			fmt.Println("Novak is elder")
		} else {
			fmt.Println("Novak and Roger are same age")
		}
	}

	// We can have 'if' without else statemnt
	if roger > novak {
		fmt.Println("Novak is Younger than Roger")
	}

	// we can write above program with 'if..else ladder'
	// Observe the else block, which is in the above specified format '} else if CONDITION {'
	if roger > novak {
		fmt.Println("Roger is elder")
	} else if roger < novak {
		fmt.Println("Novak is elder")
	} else {
		fmt.Println("Novak and Roger are same age")
	}

	// we can have statments before the condition
	// the 'number' will be availble for complete 'if..else' block
	if number := rand.Intn(11); number%2 == 0 {
		fmt.Println(number, " is Even ")
	} else {
		fmt.Println(number, " is Odd")
	}

}

/*
Program Output:
---------------

GoLang-Practise> go run .\3-Control-statements.go
Roger Age 94
Novak Age 39
Roger is elder
Novak is Younger than Roger
Roger is elder
6  is Even
GoLang-Practise>

*/
