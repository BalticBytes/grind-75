package main

func solve(nums []int, target int) []int {

	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			if target == (nums[i] + nums[j]) {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func naive(nums []int, target int) []int {

	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			if target == (nums[i] + nums[j]) {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
