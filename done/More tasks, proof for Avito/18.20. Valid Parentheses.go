package main

// Given a string s containing just the characters
// '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	runes := []rune(s) // наверно можно не перводить в руны, т..к каждая скобка 1 байт
	// но тогда мапы надо делать на байтах?
	stack := make([]rune, 0, len(runes))
	opening := map[rune]bool{'(': true, '{': true, '[': true}
	expectedClosing := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	for _, p := range runes {
		// если открывающая -> просто кладем на стек
		if opening[p] {
			stack = append(stack, p)
			continue
		}

		// получили закрывающую скобку
		// если на стеке ничего, значит ошибка
		if len(stack) == 0 {
			return false
		}

		// проверяем, что тип закрывающей для топ скобки совпадает с данной закрывающей
		top := stack[len(stack)-1]
		if expectedClosing[top] != p {
			return false
		}

		// снимаем со стека скобку
		stack = stack[:len(stack)-1]
	}

	// если на стеке есть открытые скобки, то ошибка
	if len(stack) != 0 {
		return false
	}
	return true
}
