//4. Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
package main

import (
	"fmt"
)

func Calc(arr map[string]int, s string) int {
	result := 0
	// Add first then check the previous element if it smaller than this index
	for i := 0; i < len(s); i++ {
		if i == 0 {
			result += arr[string(s[i])]
			continue
		} else {
			// Check previous element
			if arr[string(s[i-1])] < arr[string(s[i])] {
				result = result + arr[string(s[i])] - arr[string(s[i-1])]*2
			} else { //If not just Add
				result += arr[string(s[i])]
			}
		}
	}
	return result
}
func main() {

	arr := make(map[string]int)
	arr["I"] = 1
	arr["V"] = 5
	arr["X"] = 10
	arr["L"] = 50
	arr["C"] = 100
	arr["D"] = 500
	arr["M"] = 1000

	s := "X"
	fmt.Printf("%s : %d", s, Calc(arr, s))
}
