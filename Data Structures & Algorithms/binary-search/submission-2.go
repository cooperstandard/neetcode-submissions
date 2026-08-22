func search(nums []int, target int) int {
    low, high := 0, len(nums)-1
    if len(nums) == 0 || target > nums[len(nums)-1] || target < nums[0] {
        return -1
    }

    for {
        if high == low || high - low == 1 {
            if nums[high] == target {
                return high
            }
            if nums[low] == target {
                return low
            }
            return -1
        }
        pivot := (low + high) / 2
        if nums[pivot] == target {
            return pivot
        } else if nums[pivot] < target {
            low = pivot
        } else {
            high = pivot
        }
    }
}
