package array

import "testing"

func Test_arraySearch(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				array:  []int{1, 2, 4, 7, 3, 8, 9},
				target: 7,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				array:  []int{11, 1, 6, 1234, 966, 467},
				target: 966,
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				array:  []int{11, 1, 6, 1234, 966, 467},
				target: 7,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraySearch(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("arraySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
