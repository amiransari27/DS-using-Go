package recusion

import "testing"

// Helper function to build a linked list from values in order
func linkedListToSlice(head *TreeNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Right
	}
	return result
}

func TestFlattenBST(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Nil tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Only left child",
			root: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
			expected: []int{1, 2},
		},
		{
			name: "Only right child",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			expected: []int{1, 2},
		},
		{
			name: "Complete binary tree",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Larger tree (in-order traversal)",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 6},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Left-skewed tree",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Right-skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Complex multi-level tree",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 2},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 7},
					Right: &TreeNode{
						Val:   9,
						Left:  &TreeNode{Val: 10},
						Right: &TreeNode{Val: 11},
					},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 7, 8, 10, 9, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenBST(tt.root)
			got := linkedListToSlice(result)

			if len(got) != len(tt.expected) {
				t.Fatalf("flattenBST() length mismatch: got %d, want %d", len(got), len(tt.expected))
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("flattenBST() = %v; want %v", got, tt.expected)
					break
				}
			}
		})
	}
}
