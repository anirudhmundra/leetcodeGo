package linkedList

import (
	"leetcodeGo/problem/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWithElements(t *testing.T) {
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
	element := reverse(&first)
	assert.Equal(t, 3, element.Val)
}

func TestReverseWithNoElements(t *testing.T) {
	var first *model.ListNode
	element := reverse(first)
	assert.Equal(t, first, element)
}

func TestReverseWithOneElement(t *testing.T) {
	first := model.ListNode{
		Val:  1,
		Next: nil,
	}
	element := reverse(&first)
	assert.Equal(t, 1, element.Val)
}
