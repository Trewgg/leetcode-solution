// 13. Roman to Integer
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

package main

import "fmt"

func main() {

	var s string = "MCMXCIV"
	fmt.Println(romanToInt(s))

}

func romanToInt(s string) int {

	RomanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	//fmt.Println(RomanNumerals)

	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 != len(s) {
			if RomanNumerals[string(s[i])] < RomanNumerals[string(s[i+1])] {
				res -= RomanNumerals[string(s[i])]
			} else {
				res += RomanNumerals[string(s[i])]
			}
		} else {
			res += RomanNumerals[string(s[i])]
		}
		//fmt.Println(i, res)
	}

	return res

}
