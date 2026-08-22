func hammingWeight(n int) int {
    count := 0
    for i := 0 ; i < 32; i += 1 {
        count += (n >> i) & 1
    }
    return count
}
