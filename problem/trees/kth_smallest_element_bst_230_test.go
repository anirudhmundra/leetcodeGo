package trees

import (
	"leetcodeGo/problem/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthSmallestElement(t *testing.T) {
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
	element := kthSmallest(&root, 1)
	assert.Equal(t, 1, element)
}

func TestKthSmallestElementWithMoreNodes(t *testing.T) {
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
	element := kthSmallest(&root, 2)
	assert.Equal(t, 2, element)
}

func TestKthSmallestElementWithRightSubTreeCandidate(t *testing.T) {
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
	element := kthSmallest(&root, 3)
	assert.Equal(t, 3, element)
}

func TestKthSmallestElementWithEmptyRoot(t *testing.T) {
	var root model.TreeNode
	element := kthSmallest(&root, 1)
	assert.Equal(t, 0, element)
}
