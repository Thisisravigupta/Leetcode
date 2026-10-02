func reverseWords(s string) string {
    trimmed := strings.TrimSpace(s)

    i := len(trimmed) - 1
    j := i

    var ans strings.Builder

    for i >= 0 {
        for i >= 0 && trimmed[i] == ' ' {
            i--
        }

        if i < 0 {
            break
        }

        j = i

        for i >= 0 && trimmed[i] != ' ' {
            i--
        }

        if ans.Len() > 0 {
            ans.WriteByte(' ')
        } 

        ans.WriteString(trimmed[i+1:j+1])
    }

    return ans.String()
}