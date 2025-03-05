package main

// https://leetcode.com/problems/add-strings/description/
// плохо написана, много копипаста
// переделай на байты, разберись с работой строк, байтов, рун в Го
func addStrings(num1 string, num2 string) string {
	n1, n2 := []rune(num1), []rune(num2)
	resultL := max(len(n1), len(n2)) + 1
	result := make([]rune, resultL)
	i, j, k := len(n1)-1, len(n2)-1, resultL-1
	extra := 0
	for i >= 0 && j >= 0 {
		a, b := int(n1[i]-'0'), int(n2[j]-'0')
		sum := a + b + extra
		num := sum % 10
		extra = sum / 10
		result[k] = rune(num + '0') // тут можно ошибиться и просто num положить
		i--
		j--
		k--
	}

	var tail []rune
	if i >= 0 {
		tail = n1
	} else if j >= 0 {
		tail, i = n2, j
	}

	if len(tail) > 0 {
		for ; i >= 0; i-- {
			sum := int(tail[i]-'0') + extra
			num := sum % 10
			extra = sum / 10
			result[k] = rune(num + '0') // тут можно ошибиться и просто num положить
			k--
		}
	}

	if extra == 0 {
		result = result[1:]
	} else {
		result[k] = rune(extra + '0') // тут можно ошибиться и просто num положить
	}

	return string(result)
}
