func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	fi := strs[0]
	se := strs[len(strs)-1]
	cnt := 0
	for i := 0; i < len(fi); i++ {
		if fi[i] == se[i] {
			cnt++
		} else {
			break
		}
	}
	return fi[0:cnt]
}