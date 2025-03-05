package main

func reverse(x int) int {
	var x32, reversed, prev int32 = int32(x), 0, 0
	for x32 != 0 { // !!! работаем с отрицательными, поэтому != вместо >
		var lastDigit int32 = x32 % 10
		prev = reversed
		reversed = reversed*10 + lastDigit
		x32 /= 10
		if (reversed-lastDigit)/10 != prev {
			return 0
		}
	}
	return int(reversed)
}
