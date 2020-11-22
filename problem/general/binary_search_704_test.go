package general

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchWithElementFound(t *testing.T) {
	assert.Equal(t, 4, search([]int{1, 2, 3, 4, 5, 6, 7}, 5))
}

func TestBinarySearchWithNoElementFound(t *testing.T) {
	assert.Equal(t, -1, search([]int{1, 2, 3, 4, 5, 6, 7}, 9))
}
