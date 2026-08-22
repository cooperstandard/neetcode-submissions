func maxProfit(prices []int) int {
 if len(prices) == 1 {
    return 0
 }

 boughtPrice := math.MaxInt32
 soldPrice := 0
 for _, sell := range prices {
    if sell - boughtPrice > soldPrice {
        soldPrice = sell - boughtPrice
    }

    if sell < boughtPrice {
        boughtPrice = sell
    }
 }

 if soldPrice - boughtPrice < 0 {
    return 0
 }
 return soldPrice
}
