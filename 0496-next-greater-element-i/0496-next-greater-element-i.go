func nextGreaterElement(nums1 []int, nums2 []int) []int {
    res := []int{}
    for i := 0; i < len(nums1); i++ {
        idx := -1
        curr := -1
        for j := 0; j < len(nums2); j++ {
            if nums1[i] == nums2[j] {
                idx = j // 2, 0
            }
            if idx != -1 && nums1[i] < nums2[j] && idx < j {
                res = append(res, nums2[j])
                curr = 0
                break
            }
        }
        if curr == -1 {
            res = append(res, curr)
        }
    }
    return res
}