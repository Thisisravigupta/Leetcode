func zeroFilledSubarray(nums []int) int64 {
    count := 0
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            j++
            count += j
        } else {
            j = 0
        }
    }
    return int64(count)
}