package trees

import "leetcodeGo/problem/model"

func levelOrder(root *model.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	arr := []*model.TreeNode{root}
	for len(arr) > 0 {
		size := len(arr)
		tempResult := []int{}
		for size > 0 {
			temp := arr[0]
			arr = arr[1:len(arr)]
			if temp.Left != nil {
				arr = append(arr, temp.Left)
			}
			if temp.Right != nil {
				arr = append(arr, temp.Right)
			}
			tempResult = append(tempResult, temp.Val)
			size--
		}
		result = append(result, tempResult)
	}
	return result
}
