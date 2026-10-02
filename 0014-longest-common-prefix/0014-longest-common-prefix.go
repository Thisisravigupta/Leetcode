func longestCommonPrefix(strs []string) string {
    ans := strs[0]
    for i := 1; i < len(strs); i++ {
        curr := strs[i]
        temp := ""
        for j := 0; j < len(curr) && j < len(ans); j++ {
            if curr[j] == ans[j] {
                temp += string(curr[j])
            } else {
                break
            }
        }
        ans = temp
    }
    return ans
}