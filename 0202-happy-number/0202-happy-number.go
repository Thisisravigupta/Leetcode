func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	visited := make(map[int]struct{})

	for n != 1 {
		if _, exists := visited[n]; exists {
			return false
		}

		visited[n] = struct{}{}

		sum := 0
		for n > 0 {
			rem := n % 10
			sum += rem * rem
			n /= 10
		}

		n = sum
	}

	return true
}