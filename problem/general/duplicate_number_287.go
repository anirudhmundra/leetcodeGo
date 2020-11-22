package general

func findDuplicate(nums []int) int {
	numMap := map[int]struct{}{}
	for _, num := range nums {
		if _, found := numMap[num]; !found {
			numMap[num] = struct{}{}
		} else {
			return num
		}
	}
	return -1
}
