func reverse(x int) int {
    res := 0

    for x != 0 {
        lastDigit := x % 10
        x /= 10
        res = res * 10 + lastDigit

        if res < -2147483648 || res > 2147483647 {
            return 0
        }
    }
    return res
}