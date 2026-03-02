package recusion

import (
	"reflect"
	"sort"
	"testing"
)

func bruteSubsequences(s string) []string {
	var res []string
	n := len(s)
	for mask := 1; mask < (1 << n); mask++ {
		var b []byte
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				b = append(b, s[i])
			}
		}
		res = append(res, string(b))
	}
	sort.Strings(res)
	return res
}

func TestAllPossibleString(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"empty", ""},
		{"single", "a"},
		{"two distinct", "ab"},
		{"three letters", "abc"},
		{"duplicate chars", "aba"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allPossibleString(tt.in)
			exp := bruteSubsequences(tt.in)
			// treat nil and empty as equal
			if len(got) == 0 && len(exp) == 0 {
				return
			}
			if !reflect.DeepEqual(got, exp) {
				t.Fatalf("allPossibleString(%q) = %v; want %v", tt.in, got, exp)
			}
		})
	}
}

func TestAllPossibleStringProperties(t *testing.T) {
	tests := []struct {
		in string
	}{
		{"abcd"},
		{"aa"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := allPossibleString(tt.in)
			exp := bruteSubsequences(tt.in)
			if !reflect.DeepEqual(got, exp) {
				t.Fatalf("%s: got %v, want %v", tt.in, got, exp)
			}
		})
	}
}
