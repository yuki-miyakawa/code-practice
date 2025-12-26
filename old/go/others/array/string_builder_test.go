package array_test

import (
	"code-practice/others/array"
	"testing"
)

func Test_stringBuild(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			array.StringBuild()
		})
	}
}
