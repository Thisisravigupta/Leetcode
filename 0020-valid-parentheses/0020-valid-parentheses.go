func isValid(s string) bool {
    bracketMap := map[byte]byte {
        '(' : ')',
        '{' : '}',
        '[' : ']',
    }

    var stack []byte

    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
        } else {
            if len(stack) > 0 {
                top := stack[len(stack) - 1]
                if s[i] != bracketMap[top]{
                    return false
                }
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}