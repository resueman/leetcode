package main

import (
	"strconv"
)

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	signs := map[string]bool{"*": true, "+": true, "-": true, "/": true}
	for _, v := range tokens {
		if !signs[v] {
			op, _ := strconv.Atoi(v)
			stack = append(stack, op)
			continue
		}
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		result := 0
		switch v {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b
		}
		stack = append(stack, result)
	}
	return stack[0]
}
