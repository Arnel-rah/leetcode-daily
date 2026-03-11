

func removeAnagrams(words []string) []string {
    result := []string{words[0]}

    for i := 1; i < len(words); i++ {
        if !isAnagram(result[len(result)-1], words[i]) {
            result = append(result, words[i])
        }
    }

    return result
}

func isAnagram(a, b string) bool {
    sa := []rune(a)
    sb := []rune(b)

    sort.Slice(sa, func(i, j int) bool { return sa[i] < sa[j] })
    sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })

    return string(sa) == string(sb)
}