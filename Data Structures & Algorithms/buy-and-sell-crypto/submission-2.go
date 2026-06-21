func maxProfit(prices []int) int {
    b := 0
	s := 1
	maxP := 0
	for s < len(prices) {
		if prices[b] < prices[s] {
			profit := prices[s] - prices[b]
			if profit > maxP {
				maxP = profit
			}
		} else {
			b = s
		}
		s++
	}
	return maxP
}
