func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    m, n := len(num1), len(num2)
    res := make([]int, m+n)

    for i := m - 1; i >= 0; i-- {
        d1 := int(num1[i] - '0')
        for j := n - 1; j >= 0; j-- {
            d2 := int(num2[j] - '0')
            mul := d1 * d2
            p1, p2 := i+j, i+j+1
            sum := mul + res[p2]
            res[p2] = sum % 10
            res[p1] += sum / 10
        }
    }

    // Construire la chaîne
    var sb strings.Builder
    for _, v := range res {
        if sb.Len() > 0 || v > 0 {
            sb.WriteByte(byte(v + '0'))
        }
    }

    if sb.Len() == 0 {
        return "0"
    }
    return sb.String()
}