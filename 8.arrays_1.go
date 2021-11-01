package main

import "fmt"

/*
	- An array is a collection of elements of the same type. (Unlike Python list )
	- Go Arrays are fixed length, The length need to be defined at compile time.
	- Arrays are efficient, but their main limitation is their fixed size.
	- Use the Golang Slices for variable length arrays
	- Accessign array element is very fast O(1)
	- Searching for element takes time - O(n) (If not sorted)
*/

func main() {

	// Simple array
	var arr [5]int

	// we can access it using index.
	// which starts from 0
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Println(arr)
	// Output: [10 20 30 0 0]

	// Another way to define arrays
	// using the array literal '[n]Type{}' form
	myArr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(myArr)
	// Output: [10 20 30 40 50]

	// We can also '[...]' syntax, So compiler will compute the length
	arrAuto := [...]string{"Hello", "This", "is", "GoLang"}
	fmt.Println(arrAuto)
	// Output: [Hello This is GoLang]

	// Un-Initialzed Array elements will be filled with Zero's
	// GoLang doesn't have any un-initialized variables or items
	newArr := [4]int{}
	fmt.Println(newArr)
	//Output: [0 0 0 0]

	// Accessing elements of Array
	// We can access the elements of array using the Index
	// Be-careful while accessing the index, Don't go beyond array length. It will raise out of bounds error
	access := [2]int{44, 77}
	firstItem := access[0]
	secondItem := access[1]
	fmt.Println(firstItem, secondItem)
	// Output : 44 77

	/*
	 Iterate Over Array
	*/

	// 1. Using the ; 'for with range'

	iterArr := [...]string{"One", "Two", "Three"}
	for index, value := range iterArr {
		fmt.Printf("Index : %d, Value : %s \n", index, value)
	}
	/* Output:
	Index : 0, Value : One
	Index : 1, Value : Two
	Index : 2, Value : Three
	*/

	// If you don't want index, use _
	for _, value := range iterArr {
		fmt.Printf("Value : %s \n", value)
	}
	/* Output:
	Value : One
	Value : Two
	Value : Three
	*/

	// 2. C Style For loop.
	intArr := [...]int{10, 20, 30, 40, 50}
	for i := 0; i < len(intArr); i++ {
		fmt.Println(i, intArr[i])
	}
	/* Output:
	0 10
	1 20
	2 30
	3 40
	4 50
	*/

	// Iterate in Descending order
	for i := len(intArr) - 1; i >= 0; i-- {
		fmt.Println(intArr[i])
	}
	/* Output:
	50
	40
	30
	20
	10
	*/

	/*
		Built-In Functions
	*/

	// 'len' function for length
	// we also have 'cap' built-in function returns the same as the 'len' function
	fmt.Println(len(newArr)) // 4

	// Comparing Arrays
	// Both arrays need to be same length and elements and type
	a := [2]int{1, 2}
	b := [2]int{1, 2}
	if a == b {
		print("equal")
	} else {
		print("not equal")
	}

}
