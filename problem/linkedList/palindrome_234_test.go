package linkedList

import (
	"leetcodeGo/problem/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeWithOddElements(t *testing.T) {
	last := model.ListNode{
		Val:  1,
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
	assert.Equal(t, true, isPalindrome(&first))
}

func TestPalindromeWithEvenElements(t *testing.T) {
	last := model.ListNode{
		Val:  1,
		Next: nil,
	}
	third := model.ListNode{
		Val:  2,
		Next: &last,
	}
	second := model.ListNode{
		Val:  2,
		Next: &third,
	}
	first := model.ListNode{
		Val:  1,
		Next: &second,
	}
	assert.Equal(t, true, isPalindrome(&first))
}

func TestPalindromeWithOneElement(t *testing.T) {
	first := model.ListNode{
		Val:  1,
		Next: nil,
	}
	assert.Equal(t, true, isPalindrome(&first))
}

func TestPalindromeWithNoPalindromeElements(t *testing.T) {
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
	assert.Equal(t, false, isPalindrome(&first))
}

func TestPalindromeWithNoElements(t *testing.T) {
	var first *model.ListNode
	assert.Equal(t, true, isPalindrome(first))
}
