package main

func solve(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}

	best := 0
	buyIndex := 0
	sellIndex := 0

	for sellIndex < n {
		if prices[sellIndex] < prices[buyIndex] {
			buyIndex = sellIndex
		} else {
			profit := prices[sellIndex] - prices[buyIndex]
			if best < profit {
				best = profit
			}
		}

		sellIndex++
	}

	return best
}
