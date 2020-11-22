package general

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateNumber(t *testing.T) {
	assert.Equal(t, 1, findDuplicate([]int{1, 2, 3, 1}))
}

func TestDuplicateNumberWithOneElement(t *testing.T) {
	assert.Equal(t, 1, findDuplicate([]int{1, 1}))
}

func TestDuplicateNumberWithNoDuplicates(t *testing.T) {
	assert.Equal(t, -1, findDuplicate([]int{1, 2}))
}
