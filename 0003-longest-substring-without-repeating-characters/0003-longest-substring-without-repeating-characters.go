func lengthOfLongestSubstring(s string) int {
    i := 0
    res := 0
    seen := map[byte]int{}

    for j := 0; j < len(s); j++ {
        ch := s[j]

        if last, ok := seen[ch]; ok && last >= i {
            i = last + 1
        }

        seen[ch] = j
        res = max(res, j - i + 1)
    }

    return res
}