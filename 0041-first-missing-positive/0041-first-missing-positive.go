func firstMissingPositive(nums []int) int {
    n := len(nums)

    seen := map[int]bool{}

    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            seen[nums[i] - 1] = true
        }
    }

    for i := 0; i < n; i++ {
        if seen[i] == false {
            return i + 1
        }
    }

    return n + 1
}
