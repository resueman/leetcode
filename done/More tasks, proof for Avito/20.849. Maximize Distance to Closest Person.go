package main

// https://leetcode.com/problems/maximize-distance-to-closest-person/description/
func maxDistToClosest(seats []int) int {
	prevFree := -1
	freeCnt := 0
	maxDistance := -1
	for i, v := range seats {
		if v == 0 {
			if prevFree == -1 {
				prevFree = i
			}
			freeCnt++
			continue
		}

		if freeCnt == 0 {
			continue
		}

		if prevFree == 0 {
			maxDistance = i
			prevFree, freeCnt = -1, 0
			continue
		}

		if prevFree == i-1 {
			maxDistance = max(maxDistance, 1)
			prevFree, freeCnt = -1, 0
			continue
		}

		d := (i - prevFree) / 2
		if freeCnt%2 == 1 {
			d++
		}
		if d > maxDistance {
			maxDistance = d
		}
		prevFree, freeCnt = -1, 0
	}

	d := len(seats) - prevFree
	if seats[len(seats)-1] == 0 && d > maxDistance {
		return d
	}

	return maxDistance
}
