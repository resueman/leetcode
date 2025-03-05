package main

func generateParenthesis(n int) []string {
	// Закрытая дает 1 комбинацию, отрытая 2.
	// Открытых всего 3 -> 3 * 2  - 1, если нарисуешь дерево из двоек и единиц.

	result := make([]string, 0, 2*n-1)

	var genNext func(string, int, int)
	genNext = func(pattern string, totalOpened int, notClosedCnt int) {
		// In Go, the character "(" is part of the ASCII character set.
		// ASCII characters take exactly 1 byte in memory, regardless of the specific character.
		if len(pattern) == 2*n {
			result = append(result, pattern)
			return
		}
		if totalOpened < n {
			genNext(pattern+"(", totalOpened+1, notClosedCnt+1)
		}
		if totalOpened > len(pattern)-totalOpened {
			genNext(pattern+")", totalOpened, notClosedCnt-1)
		}
	}

	genNext("(", 1, 1)
	return result
}
