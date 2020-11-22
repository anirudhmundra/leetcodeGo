package general

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateWithOneElementRepeating(t *testing.T) {
	assert.Equal(t, 1, removeDuplicates([]int{1, 1, 1, 1}))
}

func TestRemoveDuplicateWithMoreThanOneElementRepeating(t *testing.T) {
	assert.Equal(t, 3, removeDuplicates([]int{1, 1, 2, 2, 3, 3, 3, 3}))
}

func TestRemoveDuplicateWithNoOneElementsRepeating(t *testing.T) {
	assert.Equal(t, 4, removeDuplicates([]int{1, 2, 3, 4}))
}

func TestRemoveDuplicateWithNoOneElements(t *testing.T) {
	assert.Equal(t, 0, removeDuplicates([]int{}))
}
