package main

func productExceptSelf(nums []int) []int {
	prefix, suffix := make([]int, len(nums)+1), make([]int, len(nums)+1)
	prefix[0], suffix[len(nums)] = 1, 1
	p_acc, s_acc := 1, 1
	for i := range nums {
		p_acc *= nums[i]
		prefix[i+1] = p_acc

		s_acc *= nums[len(nums)-i-1]
		suffix[len(nums)-1-i] = s_acc
	}

	products := make([]int, len(nums))
	for i := range nums {
		products[i] = prefix[i] * suffix[i+1]
	}
	return products
}
