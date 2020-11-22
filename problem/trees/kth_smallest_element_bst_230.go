package trees

import "leetcodeGo/problem/model"

func kthSmallest(root *model.TreeNode, k int) int {
	if root == nil {
		return 0
	}
	result := []*model.TreeNode{root}
	for true {
		for root != nil {
			result = append(result, root)
			root = root.Left
		}
		root = result[len(result)-1]
		result = result[:len(result)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}
