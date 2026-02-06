func fullJustify(words []string, maxWidth int) []string {
    result := []string{}
    n := len(words)
    i := 0

    for i < n {
        j := i
        lineLen := 0 

        for j < n && lineLen+len(words[j])+(j-i) <= maxWidth {
            lineLen += len(words[j])
            j++
        }

        numWords := j - i
        isLastLine := (j == n)

        var line string

        if isLastLine || numWords == 1 {
            line = words[i]
            for k := i + 1; k < j; k++ {
                line += " " + words[k]
            }
            line += strings.Repeat(" ", maxWidth-len(line))
        } else {
            totalChars := 0
            for k := i; k < j; k++ {
                totalChars += len(words[k])
            }
            totalSpaces := maxWidth - totalChars
            gaps := numWords - 1

            baseSpaces := totalSpaces / gaps
            extraSpaces := totalSpaces % gaps

            line = words[i]
            for k := i + 1; k < j; k++ {
                spaces := baseSpaces
                if extraSpaces > 0 {
                    spaces++
                    extraSpaces--
                }
                line += strings.Repeat(" ", spaces) + words[k]
            }
        }

        result = append(result, line)
        i = j
    }

    return result
}