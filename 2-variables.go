package main

import (
	"fmt"
	"strconv"
)

// Here We created variable 'I' with capital letter, THis case is called Pascal casing
// If your variable starts with capital letter, then you are exposing it to the Other packages as well.
// So when other packages import this package, They can use the `var I` in that package.
// This is one way GO provide the access modifiers.

/*
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.
pizza and pi do not start with a capital letter, so they are not exported.
When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
*/
var I = 100

// Here we declared var i, Which is in global scope.
// This will be visiable in main only if the `var i` not present in the main
// If `var i` present in main function, THen that local function will shadow this global variable
// Local gets the priority over global ( Like other programming Languages)
var i = 300

func main() {

	// go is statically typed language, So each variable is associated with the datatype.
	// var i int = 10

	// var i int
	// i = 10

	// go can identify the value 20 as int, so no need to specify
	var i = 20

	// You can use directly in python style, But use :=
	j := 30

	// To print the line
	fmt.Println("Values", i, j)

	// Assigning values ( As we are creating var 'k', use := )
	k := j
	fmt.Println("Assignment", k)

	// You already created the variable 'K' lets assign 'i'
	k = i // don't use the 'k := i', which gives error, as 'k' already exists
	fmt.Println("New k value :", k)

	// To print using printf
	// %v - Value and %T - Type
	fmt.Printf("Value : %v, Type : %T \n", i, i)

	// If you want to assign 'Int' to 'float32', Use the type-casting
	var m float32 = float32(i)
	fmt.Printf("Value : %v, Type : %T \n", m, m)

	// If you don't initialize the value of variable, By default value set to 0
	var v int
	fmt.Println("Default Value of int type: ", v)

	// String type
	var foo string = "Venkat"
	fmt.Println("Strings", foo)

	// Convert Integer to Ascii, use the 'strconv' package
	// If you use the direct string(i), which gives us the runes
	foo = strconv.Itoa(i)
	fmt.Println("----Type Conversion----")
	fmt.Printf("i value: %v and Type: %T \n", i, i)
	fmt.Printf("foo value: %v and Type: %T \n", foo, foo)

	// Initialing mutliple Integers
	var name1, name2 string = "Lewis", "George Russell"
	fmt.Println("Multiple vars : ", name1, name2)

	// short declaration
	// This is very useful, We can now declare the variables like we do in Python :)
	ram, laxman := 10, 20
	fmt.Printf("ram: %d, laxman:%d \n", ram, laxman)

	// constants - We can define the constants using the 'const' keyword
	// Constant values won't change throughtout the program
	// Few things like Version, Won't change often. So good idea to keep it as a constant.
	fmt.Printf("---Constants----\n")
	const version string = "2.31.4"
	fmt.Println(version)

}

/*
PROGRAM OUTPUT:

GoLang-Practise> go run .\2-variables.go
Values 20 30
Assignment 30
New k value : 20
Value : 20, Type : int
Value : 20, Type : float32
Default Value of int type:  0
Strings Venkat
----Type Conversion----
i value: 20 and Type: int
foo value: 20 and Type: string
Multiple vars :  Lewis George Russell
ram: 10, laxman:20
---Constants----
2.31.4
GoLang-Practise>
*/
