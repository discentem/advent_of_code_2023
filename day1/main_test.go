package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []line
	}{
		{
			input:    []byte("123\n456\n789"),
			expected: []line{"123", "456", "789"},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, split(test.input))
	}
}

func TestUnscramble(t *testing.T) {
	tests := []struct {
		input       []byte
		expectedSum int
	}{
		{
			input: func() []byte {
				b, err := os.ReadFile("artifacts/example_input.txt")
				assert.NoError(t, err)
				return []byte(b)
			}(),
			expectedSum: 142,
		},
		{
			input: func() []byte {
				b, err := os.ReadFile("artifacts/input.txt")
				assert.NoError(t, err)
				return []byte(b)
			}(),
			expectedSum: 54927,
		},
	}
	for _, test := range tests {
		sum, err := unscramble(test.input)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedSum, sum)
	}
}
