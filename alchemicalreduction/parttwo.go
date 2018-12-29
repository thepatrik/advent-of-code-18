package alchemicalreduction

import "unicode"

// FindShortestPolymerLen finds the shortest polymers
// length.
func FindShortestPolymerLen(s string) int {
	m := make(map[rune]int)
	runes := []rune(s)
	for _, r := range "abcdefghijklmnopqrstuvxyz" {
		res := washrune(runes, r)
		changes := 0
		for {
			res, changes = wash(res)
			if changes == 0 {
				m[r] = len(res)
				break
			}
		}
	}

	// Find the polymer with the lowest score
	_, lowest := func() (string, int) {
		lowest := len(runes)
		id := ""
		for r, score := range m {
			if score < lowest {
				lowest = score
				id = string(r)
			}
		}
		return id, lowest
	}()

	return lowest
}

func washrune(r []rune, test rune) []rune {
	buff := make([]rune, 0)

	for i := 0; i < len(r); i++ {
		if (unicode.IsLower(r[i]) && r[i] != test) || (unicode.IsUpper(r[i]) && unicode.SimpleFold(r[i]) != test) {
			buff = append(buff, r[i])
		}
	}
	return buff
}
