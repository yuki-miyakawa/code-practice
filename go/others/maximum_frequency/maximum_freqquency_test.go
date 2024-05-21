package maximumfrequency

import "testing"

func Test_maximum_frequency(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{nums: []int{1, 2, 3, 2, 1}},
			want: 4,
		},
		{
			name: "test2",
			args: args{nums: []int{1, 5, 2, 3, 6}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum_frequency(tt.args.nums); got != tt.want {
				t.Errorf("maximum_frequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
