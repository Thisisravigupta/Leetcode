func maxProfit(prices []int) int {
    buy := prices[0]
    profit := 0;
    for i := 0; i < len(prices); i++ {
        if prices[i] < buy {
            buy = prices[i]
        } else {
            profit += prices[i] - buy
            buy = prices[i]
        }
    }
    return profit
}