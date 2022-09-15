// 9. Palindrome Number
// Given an integer x, return true if x is palindrome integer.
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

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
