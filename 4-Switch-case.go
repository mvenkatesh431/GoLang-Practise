package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Random Number generator Seed
	rand.Seed(time.Now().UnixNano())

	// Generate the two random Integer values from 0-110
	var roger, novak int = rand.Intn(110), rand.Intn(100)

	fmt.Printf("Roger Age %d \n", roger)
	fmt.Printf("Novak Age %d \n", novak)

	// Switch works as you expected it to work. Matching the cases.
	// No need to have the break statement, Go breaks after the case.
	// If you specify the Expression beside the switch
	// Like here, Then you can't use the 'roger < 20' conditions at the case. This is due to the bool conversion error.
	// We also have the 'default' case, If none of the above.
	switch roger {
	case 10:
		fmt.Println("Roger is 10")
	case 30:
		fmt.Println("Roger is 10")
	case 40:
		fmt.Println("Roger is 10")
	default:
		fmt.Println("None of the above")
	}

	// But if you don't specify any expression beside the 'switch', Then you can now use The expression at the case Like below
	// switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
	switch {
	case roger <= 13:
		fmt.Println("Roger is Kid")
	case roger <= 20:
		fmt.Println("Roger is Tean")
	case roger <= 40:
		fmt.Println("Roger is Adult")
	default:
		fmt.Println("Roger is older than 40 (Does he OLDMAN!) ")
	}

	// We can have switch case with the Expressions before the "Switch Condition", Same like the 'if' block
	// The variables which are specified at the start of switch are available throughtout the switch block - here 'ageSum' is available for switch case
	switch ageSum := roger + novak; ageSum {
	case 10:
		fmt.Println("roger + novak = 10")
	// we can have multiple values at the switch case. Good one.
	case 20, 22, 30, 40:
		fmt.Println("roger + novak = is 20 or 22 or 20 or 40")
	// we can have calculations as well
	case 3 * 33:
		fmt.Println("ageJohn + agePaul = 3*33")
	default:
		fmt.Println("Nothing is matching man!")

	}

	// Example : calculate the day is weekday or weekend
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Howdy It's Weekend !!!")
	default:
		fmt.Println("Weekday, Work man!")
	}
}

/*
Program Output:

GoLang-Practise> go run .\4-Switch-case.go
Roger Age 86
Novak Age 23
None of the above
Roger is older than 40 (Does he OLDMAN!)
Nothing is matching man!
Weekday, Work man!
GoLang-Practise>

*/
