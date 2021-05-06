package ds

import (
	"reflect"
	"testing"
)

func TestOrderList(t *testing.T) {
	buffer := []int{}
	callback := func(val int) { buffer = append(buffer, val) }
	type args struct {
		vals []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantKTh int
	}{
		{
			name:    "case1",
			args:    args{vals: []int{9, 3, 7, 5, 1}, k: 1},
			want:    []int{1, 3, 5, 7, 9},
			wantKTh: 9,
		},
		{
			name:    "case1",
			args:    args{vals: []int{9, 3, 7, 5, 1}, k: 5},
			want:    []int{1, 3, 5, 7, 9},
			wantKTh: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer = []int{}

			list := NewOrderList()
			for _, val := range tt.args.vals {
				list.Add(val)
			}
			list.Traverse(callback)

			if !reflect.DeepEqual(tt.want, buffer) {
				t.Fatalf("want!=got. want:%v,got:%v", tt.want, buffer)
			}

			element, err := list.KMax(tt.args.k)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.wantKTh, element) {
				t.Fatalf("want!=got. want:%v,got:%v", tt.wantKTh, element)
			}

		})
	}
}
