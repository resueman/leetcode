package main

import "strings"

// https://leetcode.com/problems/integer-to-roman/
// сложность O(M), где M -- кол-во цифр в num. M = logN, где N = num,
// значит сложность O(logN),  где N -- данное нам число
func intToRoman(num int) string {
	romans := []string{"I", "V", "X", "L", "C", "D", "M"}
	k := 0
	reversedRoman := []string{}
	for num > 0 {
		d := num % 10
		romanD := ""
		switch { // лучше, чем else if. Просто if дает ошибку
		case d < 4:
			for i := 0; i < d; i++ {
				romanD += romans[k]
			}
		case d == 4:
			romanD = romans[k] + romans[k+1]
		case d < 9:
			romanD = romans[k+1]
			for i := 0; i < d-5; i++ {
				romanD += romans[k]
			}
		case d == 9:
			romanD = romans[k] + romans[k+2]
		}
		reversedRoman = append(reversedRoman, romanD)
		k += 2
		num /= 10
	}

	var builder strings.Builder
	for i := len(reversedRoman) - 1; i >= 0; i-- {
		builder.WriteString(reversedRoman[i])
	}

	return builder.String()
}
