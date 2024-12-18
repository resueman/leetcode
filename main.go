package main

func main() {

	// 340. Longest Substring with At Most K Distinct Characters.go
	println(longestSubstrWithAtMostKDistinctChars("", 5))                // 0
	println(longestSubstrWithAtMostKDistinctChars("r", -1))              // -1
	println(longestSubstrWithAtMostKDistinctChars("defaabaacabbd", 2))   // 5
	println(longestSubstrWithAtMostKDistinctChars("defaabaacabbd", 3))   // 9
	println(longestSubstrWithAtMostKDistinctChars("defffaabaacabbd", 4)) // 12
	println(longestSubstrWithAtMostKDistinctChars("defaabaacabbddd", 4)) // 12
	println(longestSubstrWithAtMostKDistinctChars("abcde", 1))           // 1
	println(longestSubstrWithAtMostKDistinctChars("abbcde", 1))          // 2
	println(longestSubstrWithAtMostKDistinctChars("abcde", 5))           // 5
	println(longestSubstrWithAtMostKDistinctChars("abcde", 15))          // 5
}
