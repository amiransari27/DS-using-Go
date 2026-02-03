package recusion

import (
	"slices"
	"testing"
)

func TestRatInMaze1(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]int
		n        int
		expected []string
	}{
		{
			name:     "Single cell",
			maze:     [][]int{{1}},
			n:        1,
			expected: []string{""},
		},
		{
			name: "2x2 maze blocked",
			maze: [][]int{
				{1, 0},
				{0, 1},
			},
			n:        2,
			expected: []string{},
		},
		{
			name: "2x2 maze simple path",
			maze: [][]int{
				{1, 1},
				{1, 1},
			},
			n:        2,
			expected: []string{"DR", "RD"},
		},
		{
			name: "3x3 maze with obstacles",
			maze: [][]int{
				{1, 1, 1},
				{0, 1, 0},
				{1, 1, 1},
			},
			n:        3,
			expected: []string{"RDDR"},
		},
		{
			name: "4x4 maze",
			maze: [][]int{
				{1, 1, 1, 1},
				{1, 0, 0, 1},
				{1, 0, 0, 1},
				{1, 1, 1, 1},
			},
			n: 4,
			expected: []string{
				"DDDRRR",
				"RRRDDD",
			},
		},
		{
			name: "No path exists",
			maze: [][]int{
				{1, 0},
				{0, 0},
			},
			n:        2,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a deep copy of the maze to avoid modifying the original
			mazeCopy := make([][]int, len(tt.maze))
			for i := range tt.maze {
				mazeCopy[i] = make([]int, len(tt.maze[i]))
				copy(mazeCopy[i], tt.maze[i])
			}

			got := RatInMaze1(mazeCopy, tt.n)

			// Sort both slices for comparison
			slices.Sort(got)
			slices.Sort(tt.expected)

			if !slices.Equal(got, tt.expected) {
				t.Errorf("RatInMaze1() = %v; want %v", got, tt.expected)
			}
		})
	}
}
