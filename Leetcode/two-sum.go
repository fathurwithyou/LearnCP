func twoSum(nums []int, t int) []int {
    mp := make(map[int]int)
    for i, num := range nums {
        if idx, found := mp[t-num]; found {
            return []int{idx, i}
        }
        mp[num] = i
    }
    return []int{}
}
