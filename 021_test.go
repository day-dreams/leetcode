package leetcode

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"case1", args{l1: &ListNode{Val: 1, Next: nil}, l2: nil}, nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := mergeTwoLists(tt.args.l1, tt.args.l2)
			fmt.Printf("%s\n", Json(ret))
		})
	}
}
