func isHappy(n int) bool {
    slow := n
    fast := n

    for {
        slow = getNext(slow)
        fast = getNext(getNext(fast))

        if fast == 1 {
            return true
        }

        if slow == fast {
            return false
        }
    }
}

func getNext(n int) int {
    sum := 0

    for n > 0 {
        rem := n % 10
        sum += rem * rem
        n /= 10
    }

    return sum
}
