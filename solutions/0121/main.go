package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice, maxProfit := prices[0], 0
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
