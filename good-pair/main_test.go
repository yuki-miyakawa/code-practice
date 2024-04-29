package main

import "testing"

func Test_goodPair(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{array: []int{1, 2, 3, 1, 1, 3}},
			want: 4,
		},
		{
			name: "test2",
			args: args{array: []int{1, 1, 1, 1}},
			want: 6,
		},
		{
			name: "test3",
			args: args{array: []int{1, 2, 3}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodPair(tt.args.array); got != tt.want {
				t.Errorf("goodPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
