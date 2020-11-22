package linkedList

import (
	"leetcodeGo/problem/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNodeWithOddNumberOfElements(t *testing.T) {
	last := model.ListNode{
		Val:  3,
		Next: nil,
	}
	second := model.ListNode{
		Val:  2,
		Next: &last,
	}
	first := model.ListNode{
		Val:  1,
		Next: &second,
	}
	element := middleNode(&first)
	assert.Equal(t, 2, element.Val)
}

func TestMiddleNodeWithEvenNumberOfElements(t *testing.T) {
	second := model.ListNode{
		Val:  2,
		Next: nil,
	}
	first := model.ListNode{
		Val:  1,
		Next: &second,
	}
	element := middleNode(&first)
	assert.Equal(t, 2, element.Val)
}

func TestMiddleNodeWithOneElement(t *testing.T) {
	first := model.ListNode{
		Val:  1,
		Next: nil,
	}
	element := middleNode(&first)
	assert.Equal(t, 1, element.Val)
}

func TestMiddleNodeWithNoElement(t *testing.T) {
	var first *model.ListNode
	element := middleNode(first)
	assert.Equal(t, first, element)
}
