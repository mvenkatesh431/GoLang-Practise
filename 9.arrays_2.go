package main

import (
	"fmt"
)

var newValue string = "NewValue"

func updateElementByValue(arr [1]string) {
	// arr is copy, even if you update it won't effect the original array
	arr[0] = newValue
}

func updateElementByPointer(arr *[1]string) {
	// arr is copy, even if you update it won't effect the original array
	arr[0] = newValue
}

func main() {

	// Passing Array to function
	// Two ways
	// Passing by value - creates the copy at function
	// Passing by Pointer - Same array will be used at function, all changes will reflect on calling function

	arr := [1]string{"OldValue"}
	//updateElementByValue(arr)

	// Pass the array by address
	updateElementByPointer(&arr)
	if arr[0] == "OldValue" {
		fmt.Println("Value is not updated")
	} else {
		fmt.Println("Value is Updated")
	}

	// Assigning array to another array creates a Deep copy
	// This creates the deep copy
	newArray := arr
	newArray[0] = "One"
	fmt.Println(arr[0], newArray[0])
	// Output: NewValue One

	// To Create the shallow copy Use the address of Array
	newArray1 := &arr
	// Both 'arr' and 'newArray1' index 0 value will be modified.
	newArray1[0] = "two"
	fmt.Println(arr[0], newArray1[0])
	// Output:  two two

	// Two dimension arrays
	// Empty strings are 'nil'
	twoArray := [3][2]string{}
	twoArray[0][1] = "0-1"
	twoArray[1][1] = "1-1"
	twoArray[2][0] = "2-0"
	fmt.Println(twoArray)
	// Output:
	// [[ 0-1] [ 1-1] [2-0 ]]

}
