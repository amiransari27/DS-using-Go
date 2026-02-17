package recusion

import "testing"

// Helper function to create a linked list node
func createNode(val int) *LinkList1 {
	return &LinkList1{
		Val:    val,
		Next:   nil,
		Bottom: nil,
	}
}

// Helper function to convert flattened list to slice for easy comparison
func flattenListToSlice(head *LinkList1) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Bottom
	}
	return result
}

func TestFlattenLinkList1(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *LinkList1
		expected []int
	}{
		{
			name: "Single node",
			setup: func() *LinkList1 {
				return createNode(1)
			},
			expected: []int{1},
		},
		{
			name: "Two nodes in horizontal list",
			setup: func() *LinkList1 {
				head := createNode(1)
				head.Next = createNode(2)
				return head
			},
			expected: []int{1, 2},
		},
		{
			name: "Two nodes in vertical list",
			setup: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(2)
				return head
			},
			expected: []int{1, 2},
		},
		{
			name: "Simple 2x2 matrix",
			setup: func() *LinkList1 {
				// Node 1: 1->2 (horizontal), 3 (vertical)
				// Node 2: 2->4 (horizontal), 5 (vertical)
				h1 := createNode(1)
				h1.Next = createNode(2)
				h1.Bottom = createNode(3)

				h1.Next.Bottom = createNode(4)
				h1.Bottom.Bottom = createNode(5)

				return h1
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Three horizontal nodes each with bottom",
			setup: func() *LinkList1 {
				// Row 1: 1->2->3
				//        |  |  |
				//        4  5  6
				h1 := createNode(1)
				h2 := createNode(2)
				h3 := createNode(3)

				h1.Next = h2
				h2.Next = h3

				h1.Bottom = createNode(4)
				h2.Bottom = createNode(5)
				h3.Bottom = createNode(6)

				return h1
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Uneven depths",
			setup: func() *LinkList1 {
				// 1->2->3
				// |     |
				// 4     5->6
				h1 := createNode(1)
				h2 := createNode(2)
				h3 := createNode(3)

				h1.Next = h2
				h2.Next = h3

				h1.Bottom = createNode(4)
				h3.Bottom = createNode(5)
				h3.Bottom.Bottom = createNode(6)

				return h1
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Complex structure with varied values",
			setup: func() *LinkList1 {
				// 5->10->19->28
				// |   |   |    |
				// 7  20  22   29
				// |       |
				// 8      23->24
				//             |
				//            26
				h1 := createNode(5)
				h2 := createNode(10)
				h3 := createNode(19)
				h4 := createNode(28)

				h1.Next = h2
				h2.Next = h3
				h3.Next = h4

				h1.Bottom = createNode(7)
				h1.Bottom.Bottom = createNode(8)

				h2.Bottom = createNode(20)

				h3.Bottom = createNode(22)
				h3.Bottom.Bottom = createNode(23)
				h3.Bottom.Bottom.Bottom = createNode(24)
				h3.Bottom.Bottom.Bottom.Bottom = createNode(26)

				h4.Bottom = createNode(29)

				return h1
			},
			expected: []int{5, 7, 8, 10, 19, 20, 22, 23, 24, 26, 28, 29},
		},
		{
			name: "All nodes in vertical only",
			setup: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(2)
				head.Bottom.Bottom = createNode(3)
				head.Bottom.Bottom.Bottom = createNode(4)
				return head
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "Empty list (nil)",
			setup: func() *LinkList1 {
				return nil
			},
			expected: []int{},
		},
		{
			name: "Duplicate values",
			setup: func() *LinkList1 {
				h1 := createNode(5)
				h2 := createNode(5)
				h1.Next = h2
				h1.Bottom = createNode(5)
				h2.Bottom = createNode(5)
				return h1
			},
			expected: []int{5, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := tt.setup()
			result := flattenLinkList1(head)
			got := flattenListToSlice(result)

			if len(got) != len(tt.expected) {
				t.Fatalf("Length mismatch: got %d elements, want %d", len(got), len(tt.expected))
			}

			for i, v := range got {
				if v != tt.expected[i] {
					t.Fatalf("Value mismatch at index %d: got %d, want %d", i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestMergeFlattenLinkList1_2(t *testing.T) {
	tests := []struct {
		name     string
		setupH1  func() *LinkList1
		setupH2  func() *LinkList1
		expected []int
	}{
		{
			name: "Both nil",
			setupH1: func() *LinkList1 {
				return nil
			},
			setupH2: func() *LinkList1 {
				return nil
			},
			expected: []int{},
		},
		{
			name: "First nil",
			setupH1: func() *LinkList1 {
				return nil
			},
			setupH2: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(2)
				return head
			},
			expected: []int{1, 2},
		},
		{
			name: "Second nil",
			setupH1: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(2)
				return head
			},
			setupH2: func() *LinkList1 {
				return nil
			},
			expected: []int{1, 2},
		},
		{
			name: "Single node each",
			setupH1: func() *LinkList1 {
				return createNode(3)
			},
			setupH2: func() *LinkList1 {
				return createNode(1)
			},
			expected: []int{1, 3},
		},
		{
			name: "Two nodes each",
			setupH1: func() *LinkList1 {
				head := createNode(2)
				head.Bottom = createNode(5)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(3)
				return head
			},
			expected: []int{1, 2, 3, 5},
		},
		{
			name: "Three nodes each",
			setupH1: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(4)
				head.Bottom.Bottom = createNode(7)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(2)
				head.Bottom = createNode(3)
				head.Bottom.Bottom = createNode(5)
				return head
			},
			expected: []int{1, 2, 3, 4, 5, 7},
		},
		{
			name: "Duplicate values",
			setupH1: func() *LinkList1 {
				head := createNode(2)
				head.Bottom = createNode(2)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(2)
				head.Bottom = createNode(3)
				return head
			},
			expected: []int{2, 2, 2, 3},
		},
		{
			name: "Negative numbers",
			setupH1: func() *LinkList1 {
				head := createNode(-5)
				head.Bottom = createNode(0)
				head.Bottom.Bottom = createNode(5)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(-10)
				head.Bottom = createNode(3)
				head.Bottom.Bottom = createNode(10)
				return head
			},
			expected: []int{-10, -5, 0, 3, 5, 10},
		},
		{
			name: "All same values",
			setupH1: func() *LinkList1 {
				head := createNode(5)
				head.Bottom = createNode(5)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(5)
				head.Bottom = createNode(5)
				return head
			},
			expected: []int{5, 5, 5, 5},
		},
		{
			name: "First list longer",
			setupH1: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(3)
				head.Bottom.Bottom = createNode(8)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(2)
				head.Bottom = createNode(4)
				head.Bottom.Bottom = createNode(6)
				return head
			},
			expected: []int{1, 2, 3, 4, 6, 8},
		},
		{
			name: "Second list longer",
			setupH1: func() *LinkList1 {
				head := createNode(2)
				head.Bottom = createNode(4)
				head.Bottom.Bottom = createNode(6)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(3)
				head.Bottom.Bottom = createNode(5)
				head.Bottom.Bottom.Bottom = createNode(7)
				return head
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "First completely smaller",
			setupH1: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(2)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(10)
				head.Bottom = createNode(20)
				return head
			},
			expected: []int{1, 2, 10, 20},
		},
		{
			name: "Complex interleaved",
			setupH1: func() *LinkList1 {
				head := createNode(1)
				head.Bottom = createNode(4)
				head.Bottom.Bottom = createNode(7)
				head.Bottom.Bottom.Bottom = createNode(9)
				return head
			},
			setupH2: func() *LinkList1 {
				head := createNode(2)
				head.Bottom = createNode(3)
				head.Bottom.Bottom = createNode(5)
				head.Bottom.Bottom.Bottom = createNode(6)
				head.Bottom.Bottom.Bottom.Bottom = createNode(8)
				return head
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h1 := tt.setupH1()
			h2 := tt.setupH2()
			result := mergeFlattenLinkList1_2(h1, h2)
			got := flattenListToSlice(result)

			if len(got) != len(tt.expected) {
				t.Fatalf("Length mismatch: got %d elements, want %d; got values %v", len(got), len(tt.expected), got)
			}

			for i, v := range got {
				if v != tt.expected[i] {
					t.Fatalf("Value mismatch at index %d: got %d, want %d", i, v, tt.expected[i])
				}
			}
		})
	}
}
