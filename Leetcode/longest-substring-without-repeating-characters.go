func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func lengthOfLongestSubstring(s string) int {
    ans := 0
    mp := make(map[rune]int) 
    start := 0              
    for i, char := range s {
        if lastIdx, ok := mp[char]; ok && lastIdx >= start {
            start = lastIdx + 1 
        }
        mp[char] = i              
        ans = max(ans, i-start+1) 
    }
    return ans
}
