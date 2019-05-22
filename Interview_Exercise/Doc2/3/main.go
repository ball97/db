//3. Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
package main

import "fmt"

func ReverseInt(num int) bool {
	if num < 0 {
		return false
	}
	temp := num
	result := 0
	for {
		remainder := num % 10
		result = (result*10 + remainder)
		num = num / 10
		if num < 1 {
			break
		}
	}
	if result == temp {
		return true
	} else {
		return false
	}

}

func main() {
	num := -12
	result := ReverseInt(num)
	if result {
		fmt.Printf("This interger is a panlindrome")
	} else {
		fmt.Printf("This interger is not a panlindrome")
	}
}
