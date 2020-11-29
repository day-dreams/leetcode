package leetcode

import (
	"testing"
)

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{A: []int{1, 2, 1}},
			want: 0,
		},
		{
			args: args{A: []int{2, 1, 2}},
			want: 5,
		},
		{
			args: args{A: []int{3, 2, 3, 4}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.A); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
