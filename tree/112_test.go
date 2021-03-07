package tree

import (
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "3",
			args: args{
				targetSum: 3,
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -2,
						Left: &TreeNode{
							Val:  1,
							Left: &TreeNode{Val: -1},
						},
						Right: &TreeNode{Val: 3}},
					Right: &TreeNode{
						Val: -3,
						Left: &TreeNode{
							Val: -2,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
