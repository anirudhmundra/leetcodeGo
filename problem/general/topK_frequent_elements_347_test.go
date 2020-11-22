package general

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequentElements(t *testing.T) {
	assert.Equal(t, []int{3, 1}, topKFrequent([]int{1, 1, 1, 1, 3, 3, 2}, 2))
}

func TestTopKFrequentElementsWithOneElement(t *testing.T) {
	assert.Equal(t, []int{1}, topKFrequent([]int{1, 1, 1, 1}, 1))
}
