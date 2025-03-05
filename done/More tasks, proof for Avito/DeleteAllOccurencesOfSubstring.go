package main

import (
	"strings"
)

func DeleteFromString(s string) string {
	sub := "?%#"
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub { // аккуратно, можно выйти за границы массива
			i += len(sub) - 1 // если идешь в цикле for и меняешь i, то помни, что перед следующей итерацией i увеличится
		} else {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}

func DeleteFromStringInPlace(s string) string {
	sub := []rune("?%#")
	runes := []rune(s)
	l := 0
	for r := 0; r < len(runes); r++ {
		if r+len(sub) > len(runes) { // проверка выхода за границы прежде чем взять срез
			runes[l] = runes[r]
			l++
			continue
		}

		if areEqual(sub, runes[r:r+len(sub)]) {
			r += len(sub) - 1 // если идешь в цикле for и меняешь i, то помни, что перед следующей итерацией i увеличится
		} else {
			runes[l] = runes[r]
			l++
		}
	}
	runes = runes[:l]
	return string(runes)
}

func areEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
