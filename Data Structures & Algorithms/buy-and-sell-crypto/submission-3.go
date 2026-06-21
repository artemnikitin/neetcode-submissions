func maxProfit(prices []int) int {
    maxP := 0
	minBuy := 1000
	for _, v := range prices {
		if v-minBuy > maxP {
			maxP = v - minBuy
		}
		if v < minBuy {
			minBuy = v
		}
	}
	return maxP
}
