// 12. Integer to Roman
// Input: num = 1994
// Output: "MCMXCIV"
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

package main

import "fmt"

func main() {

	var n int = 1994
	fmt.Println(intToRoman(n))

}

func intToRoman(num int) string {

	RomanNumerals := map[string]int{
		"I":   1,
		"II":  2,
		"III": 3,
		"IV":  4,
		"V":   5,
		"IX":  9,
		"X":   10,
		"XL":  40,
		"L":   50,
		"XC":  90,
		"C":   100,
		"CD":  400,
		"D":   500,
		"CM":  900,
		"M":   1000,
	}
	//fmt.Println(RomanNumerals)

	var m []int
	var temp int
	var res string

	for num > temp {
		//fmt.Println(num, temp)
		num -= temp
		temp = tempNum(RomanNumerals, num)
		m = append(m, temp)

	}

	for i := range m {
		for key, value := range RomanNumerals {
			if value == m[i] {
				//fmt.Println(key, value)
				res += key
			}
		}
	}

	fmt.Println(m)

	return res
}

func tempNum(RomanNumerals map[string]int, num int) int {

	temp := 0
	for _, value := range RomanNumerals {
		if num >= value {
			temp = max(temp, value)
		}
	}

	return temp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
