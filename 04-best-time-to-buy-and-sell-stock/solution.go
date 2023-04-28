package main

func solve(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	best := 0

	for i := range prices {
		for j := i; j < len(prices); j++ {
			profit := (prices[j] - prices[i])
			if best < profit {
				best = profit
			}
		}
	}

	return best
}
