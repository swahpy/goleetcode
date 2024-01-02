package twosum

func twoSum(nums []int, target int) []int {
	nummap := make(map[int]int)
	res := make([]int, 2)
	for i, v := range nums {
		ano := target - v
		if index, ok := nummap[ano]; ok {
			if index != i {
				res[0] = index
				res[1] = i
				break
			}
		}
		nummap[v] = i
	}
	return res
}
