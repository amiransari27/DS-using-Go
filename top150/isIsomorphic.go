package top150

import "strings"

// e g g
// a d d
func IsIsomorphic(s string, t string) bool {
	s = strings.Trim(s, "\u0000")
	t = strings.Trim(t, "\u0000")
	n := len(s)

	if len(s) != len(t) {
		return false
	}
	map_ := make(map[byte]byte)
	_map := make(map[byte]byte)

	for i := 0; i < n; i++ {
		chr1 := s[i]
		chr2 := t[i]

		if (map_[chr1] != 0 && map_[chr1] != chr2) ||
			(_map[chr2] != 0 && _map[chr2] != chr1) {
			return false
		}

		map_[chr1] = chr2
		_map[chr2] = chr1
	}

	return true
}
