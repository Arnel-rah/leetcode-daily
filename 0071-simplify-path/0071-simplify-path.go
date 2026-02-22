func simplifyPath(path string) string {
    stack := []string{}
    
    for _, part := range strings.Split(path, "/") {
        if part == "" || part == "." {
            continue
        }
        if part == ".." {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        } else {
            stack = append(stack, part)
        }
    }
    
    var sb strings.Builder
    sb.WriteByte('/')
    for i, dir := range stack {
        if i > 0 {
            sb.WriteByte('/')
        }
        sb.WriteString(dir)
    }
    
    return sb.String()
}