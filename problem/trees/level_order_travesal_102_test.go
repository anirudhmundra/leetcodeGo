package trees

import (
	"leetcodeGo/problem/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrderTraversal(t *testing.T) {
	third := model.TreeNode{
		Val: 3,
	}
	first := model.TreeNode{
		Val: 1,
	}
	root := model.TreeNode{
		Val:   2,
		Left:  &first,
		Right: &third,
	}
	elements := levelOrder(&root)
	assert.Equal(t, [][]int{{2}, {1, 3}}, elements)
}

func TestLevelOrderTraversalWithMoreNodes(t *testing.T) {
	fifth := model.TreeNode{
		Val: 5,
	}
	third := model.TreeNode{
		Val: 3,
	}
	first := model.TreeNode{
		Val: 1,
	}
	second := model.TreeNode{
		Val:   2,
		Right: &third,
		Left:  &first,
	}
	root := model.TreeNode{
		Val:   4,
		Left:  &second,
		Right: &fifth,
	}
	elements := levelOrder(&root)
	assert.Equal(t, [][]int{{4}, {2, 5}, {1, 3}}, elements)
}
