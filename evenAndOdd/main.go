// Create a new package
package main

import "fmt"

func main() {
	// In the main function, create a slice of nts form 0 through 10
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Iterate through the slice, For each element, chek to see if the number is even or odd
	for _, num := range nums {
		// If the value is even print out "even", if it is odd print out "odd"
		if num%2 == 0 {
			fmt.Println(num, "Even")
		} else {
			fmt.Println(num, "Odd")
		}
	}
}
