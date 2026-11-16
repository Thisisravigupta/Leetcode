func nextGreaterElement(nums1 []int, nums2 []int) []int {
    res := []int{}
    for i := 0; i < len(nums1); i++ {
        idx := -1
        nge := -1
        for j := 0; j < len(nums2); j++ {
            if nums1[i] == nums2[j] {
                idx = j 
            }
            if idx != -1 && nums1[i] < nums2[j] && idx < j {
                nge = nums2[j]
                break
            }
        }
        res = append(res, nge)
    }
    return res
}