package ds

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree_Traverse(t *testing.T) {

	buffer := []int{}
	callback := func(val int) { buffer = append(buffer, val) }
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{vals: []int{9, 3, 7, 5, 1}},
			want: []int{1, 3, 5, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer = []int{}

			tree := NewBinarySearchTree()
			for _, val := range tt.args.vals {
				tree.Add(val)
			}
			tree.Traverse(callback)

			if !reflect.DeepEqual(tt.want, buffer) {
				t.Fatalf("want!=got. want:%v,got:%v", tt.want, buffer)
			}
		})
	}
}
