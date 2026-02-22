package repl

import (
	"testing"

	// "github.com/goforj/godump"
	"github.com/stretchr/testify/require"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello            world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "           world",
			expected: []string{"world"},
		},
	}

	for _, c := range cases {
		res := cleanUpInput(c.input)
		require.Equal(t, c.expected, res, "case failed %s", c.input)
	}
}
