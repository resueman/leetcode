package main

// Дается строка, состоящая из букв X, Y, O.
// Найти минимальное расстояние между X и Y (в межбуквенных интервалах).
// Если в строке обеих сразу букв нет, то вернуть 0.
// "XX" -> 0
// "YY" ->0
// "XY" -> 1
// "XOY" -> 2
// "OOOXOOOYOXOO" -> 2
// "OOXXOOY" -> 3

func MinDistanceBetweenXAndY(s string) int {
	prevX, prevY := -1, -1
	minD := len(s) + 1 // !!!
	for r := 0; r < len(s); r++ {
		if s[r] == 'X' {
			if prevY != -1 {
				minD = min(r-prevY, minD)
			}
			prevX = r
		} else if s[r] == 'Y' {
			if prevX != -1 {
				minD = min(r-prevX, minD)
			}
			prevY = r
		}
	}

	if prevX == -1 || prevY == -1 { // !!! ну или minD == len(s) + 1
		return 0
	}

	return minD
}

func MinDistanceBetweenXAndY2(s string) int {
	prevX, prevY := -1, -1
	prevChar := byte('O')
	minD := len(s) + 1 // !!!
	for r := 0; r < len(s); r++ {
		if s[r] != 'X' && s[r] != 'Y' {
			continue
		}

		if s[r] == 'X' {
			if prevChar == 'Y' {
				minD = min(r-prevY, minD)
			}
			prevChar, prevX = 'X', r
			continue
		}

		// s[r] == 'Y'
		if prevChar == 'X' {
			minD = min(r-prevX, minD)
		}
		prevChar, prevY = 'Y', r
	}

	if minD == len(s)+1 { // !!!
		return 0
	}

	return minD
}
