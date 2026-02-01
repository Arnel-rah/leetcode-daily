function lengthOfLongestSubstring(s: string): number {
    if (s.length <= 1) return s.length;

    const seen = new Set<string>();
    let left = 0;
    let maxLength = 0;

    for (let right = 0; right < s.length; right++) {
        const char = s[right];
        while (seen.has(char)) {
            seen.delete(s[left]);
            left++;
        }
        seen.add(char);
        maxLength = Math.max(maxLength, right - left + 1);
    }

    return maxLength;
}