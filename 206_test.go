package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {

				for got != nil {
					t.Logf("%v ", got.Val)
					got = got.Next
				}

				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
