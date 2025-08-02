package array

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		array []int
		index int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				array: []int{1, 2, 3, 4, 5},
				index: 2,
				value: 10,
			},
			want: []int{1, 2, 10, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.array, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert2(t *testing.T) {
	type args struct {
		array []int
		index int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				array: []int{1, 2, 3, 4, 5},
				index: 2,
				value: 10,
			},
			want: []int{1, 2, 10, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert2(tt.args.array, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert2() = %v, want %v", got, tt.want)
			}
		})
	}
}
