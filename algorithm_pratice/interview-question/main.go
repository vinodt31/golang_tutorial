package main

import (
	"fmt"
	"slices"
)

func convertStringToArray() []rune {
	name := "Alice"
	convertStringToArray := []rune(name)
	//Go will print the numeric ASCII/Unicode values, not the letters:
	fmt.Println("The name as an array of characters:", convertStringToArray) // [65 108 105 99 101]

	//To see them as [A l i c e], you must use fmt.Printf("%c", convertStringToArray) as shown in the examples above.
	fmt.Printf("%c\n", convertStringToArray)

	return convertStringToArray
}

func isPalindrome(name string) bool {
	str := []rune(name)
	left, right := 0, len(name)-1

	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}
	return true
}

func reverseArray(arr []string) {
	left := 0
	right := len(arr) - 1

	for left < right {
		// Swap the elements in-place
		arr[left], arr[right] = arr[right], arr[left]

		// Move the pointers toward the center
		left++
		right--
	}
}

func sortArrayValue() {
	// sort array of integers
	numbers := []int{5, 2, 9, 1, 7}
	// Sorts the slice in place in ascending order
	//slices.Sort(numbers)
	//fmt.Println(numbers) // Output: [1 2 5 7 9]
	// sort array of integer in descending order
	slices.SortFunc(numbers, func(a, b int) int {
		return b - a
	})

	fmt.Println(numbers)

	// sort array of strings
	letters := []string{"v", "i", "n", "o", "d"}
	// Sorts the slice in place in ascending order
	slices.Sort(letters)
	fmt.Println(letters) // Output: [d i n o v]
}

func main() {
	// fmt.Println(isPalindrome("racecar"))  // true

	/*
		letters := []string{"v", "i", "n", "o", "d"}
		reverseArray(letters)
		fmt.Println(letters) // Output: [d o n i v]
	*/

	//sortArrayValue()

	for i := 0; i < 4; i++ {
		i := i // "Tricks" the compiler into creating a block-scoped variable
		go func() {
			fmt.Println(i)
		}()
	}
}
