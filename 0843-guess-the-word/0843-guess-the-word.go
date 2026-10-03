/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
func findSecretWord(words []string, master *Master) {
    candidates := append([]string{}, words...)

    for len(candidates) > 0 {

        // Choose a guess that minimizes the worst-case remaining candidates
        guess := chooseGuess(candidates)

        matches := master.Guess(guess)

        if matches == 6 {
            return
        }

        filtered := []string{}

        for _, word := range candidates {
            if match(word, guess) == matches {
                filtered = append(filtered, word)
            }
        }

        candidates = filtered
    }
}

func chooseGuess(candidates []string) string {
    bestGuess := candidates[0]
    bestScore := len(candidates)

    for _, guess := range candidates {
        groups := make([]int, 7)

        for _, word := range candidates {
            matches := match(word, guess)
            groups[matches]++
        }

        maxGroup := 0

        for _, count := range groups {
            if count > maxGroup {
                maxGroup = count
            }
        }

        if maxGroup < bestScore {
            bestScore = maxGroup
            bestGuess = guess
        }
    }

    return bestGuess
}

func match(a, b string) int {
    count := 0

    for i := 0; i < 6; i++ {
        if a[i] == b[i] {
            count++
        }
    }

    return count
}