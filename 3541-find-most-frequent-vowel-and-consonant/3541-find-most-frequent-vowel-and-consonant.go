func maxFreqSum(s string) int {
    freq := make(map[rune]int)

    for _, ch := range s {
        freq[ch]++
    }

    maxVowel := 0
    maxConsonant := 0

    for ch, count := range freq {
        switch ch {
        case 'a', 'e', 'i', 'o', 'u':
            if count > maxVowel {
                maxVowel = count
            }
        default:
            if count > maxConsonant {
                maxConsonant = count
            }
        }
    }

    return maxVowel + maxConsonant
}
