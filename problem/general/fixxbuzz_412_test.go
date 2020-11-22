package general

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzForOddInput(t *testing.T) {
	result := []string{
		"1", "2", "Fizz",
		"4", "Buzz", "Fizz",
		"7", "8", "Fizz",
		"Buzz", "11", "Fizz",
		"13", "14", "FizzBuzz",
	}
	assert.Equal(t, result, fizzBuzz(15))
}

func TestFizzBuzzForEvenInput(t *testing.T) {
	result := []string{
		"1", "2", "Fizz",
		"4", "Buzz", "Fizz",
	}
	assert.Equal(t, result, fizzBuzz(6))
}

func TestFizzBuzzForZeroInput(t *testing.T) {
	result := []string{}
	assert.Equal(t, result, fizzBuzz(0))
}
