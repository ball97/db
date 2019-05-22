//2. Given a 32-bit signed integer, reverse digits of an integer.
package main

import (
	"fmt"
)

func ReverseInt(num int) int {
	if num < 0 {
		num = num * -1
	}
	result := 0
	for {
		remainder := num % 10
		result = (result*10 + remainder)
		num = num / 10
		if num < 1 {
			break
		}
	}
	return result * -1
}

func main() {
	num := 92090
	result := ReverseInt(num)
	fmt.Printf("result: %d", result)
}
