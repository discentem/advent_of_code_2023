package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	}
	for _, test := range tests {
		assert.Equal(t, test.expectedSum, unscramble(test.input))
	}
}
