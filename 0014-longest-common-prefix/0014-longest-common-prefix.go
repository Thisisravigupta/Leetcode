func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    ans := strs[0]

    for i := 1; i < len(strs); i++ {
        curr := strs[i]
        j := 0
        for j < len(curr) && j < len(ans) {
            if curr[j] != ans[j] {
                break
            }
            j++
        }
        ans = ans[:j]

        if ans == "" {
            return ""
        }
    }
    return ans
}