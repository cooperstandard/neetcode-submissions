func maxSubArray(nums []int) int {
   sum := 0
   maxSum := nums[0]
    for _, v := range nums {
        sum = sum + v
        if sum < v {
            sum = v
        }
        if sum > maxSum {
            maxSum = sum
        }
    }
    return maxSum
}
