func countBits(n int) []int {
    // TODO: implement
    ans := make([]int, n+1)
    for i := 0; i <= n; i++ {
        ans[i] = countOnes(i)
    }
    return ans
}

func countOnes(n int) int {
    count := 0
    for n > 0 {
        count += n & 1
        n >>= 1
    }
    return count
}
