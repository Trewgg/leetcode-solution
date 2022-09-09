// 9. Palindrome Number
// Given an integer x, return true if x is palindrome integer.
// An integer is a palindrome when it reads the same backward as forward.
// For example, 121 is a palindrome while 123 is not.

package main

import "fmt"

func main() {

	x := 121
	fmt.Println(x)
	fmt.Println(isPalindrome(x))

}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var n, a int
	n = x

	for x > 0 {
		a = a*10 + x%10
		x = x / 10
	}

	return n == a

}
