package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPrintWords(t *testing.T) {
	type args struct {
		words      []string
		maxLineLen int
	}
	tests := []struct {
		name     string
		args     args
		expected []string
	}{
		{
			name: "success",
			args: args{
				words:      []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxLineLen: 16,
			},
			expected: []string{"This    is    an", "example of  text", "justification."},
		},
		{
			name: "test2",
			args: args{
				words:      []string{"This", "is", "an", "example", "of", "text", "justification.", "t"},
				maxLineLen: 16,
			},
			expected: []string{"This    is    an", "example of  text", "justification. t"},
		},
		{
			name: "test3",
			args: args{
				words:      []string{"This"},
				maxLineLen: 16,
			},
			expected: []string{"This"},
		},
		{
			name: "test4",
			args: args{
				words:      []string{"This", "is", "test4"},
				maxLineLen: 16,
			},
			expected: []string{"This is  test4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.DeepEqual(t, PrintWords(tt.args.words, tt.args.maxLineLen), tt.expected)
		})
	}
}
