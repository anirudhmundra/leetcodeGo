package general

import "sort"

func topKFrequent(nums []int, k int) []int {
	resultMap := map[int]int{}
	for _, num := range nums {
		if count, ok := resultMap[num]; ok {
			resultMap[num] = count + 1
		} else {
			resultMap[num] = 1
		}
	}
	type element struct {
		Key   int
		Count int
	}
	elementArr := []element{}
	for key, val := range resultMap {
		elementArr = append(elementArr, element{Key: key, Count: val})
	}
	sort.Slice(elementArr, func(i, j int) bool {
		return elementArr[i].Count > elementArr[j].Count
	})
	resultArr := []int{}
	for ; k > 0; k-- {
		resultArr = append(resultArr, elementArr[k-1].Key)
	}
	return resultArr
}
