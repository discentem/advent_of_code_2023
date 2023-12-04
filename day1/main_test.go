package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOnlyNums(t *testing.T) {
	tests := []struct {
		input    line
		expected line
	}{
		{
			input:    "a1c2b3",
			expected: "123",
		},
		{
			input:    "a1c2b3zzzz",
			expected: "123",
		},
		{
			input:    "abc1234567890abc",
			expected: "1234567890",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, test.input.onlyNums())
	}
}

func TestNumFromIndices(t *testing.T) {
	tests := []struct {
		input       line
		indicies    []int
		expected    int
		expectedErr func(error) bool
	}{
		{
			input:    "1234567890",
			indicies: []int{0, 8},
			expected: 19,
			expectedErr: func(err error) bool {
				return err == nil
			},
		},
		{
			input:    "1234567890",
			indicies: []int{0, 15},
			expected: 0,
			expectedErr: func(err error) bool {
				return err.Error() == "invalid index 15"
			},
		},
	}
	for _, test := range tests {
		actual, err := test.input.numFromIndicies(test.indicies)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if test.expectedErr != nil {
			assert.Equal(t, test.expectedErr(err), true)
			continue
		}
		require.NotNil(t, actual)
		assert.Equal(t, test.expected, *actual)
	}
}

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
	}
	for _, test := range tests {
		assert.Equal(t, test.expectedSum, unscramble(test.input))
	}
}
