package alchemicalreduction

import (
	"sync"
	"unicode"
)

const (
	alpha = "abcdefghijklmnopqrstuvxyz"
)

// FindShortestPolymerLen finds the shortest polymers
// length.
func FindShortestPolymerLen(s string) int {
	m := make(map[rune]int)
	runes := []rune(s)

	var wg sync.WaitGroup
	wg.Add(len(alpha))

	for _, r := range alpha {
		// Run each rune in a separate goroutine
		go func(alpha rune) {
			defer wg.Done()

			res := washrune(runes, alpha)
			changes := 0
			for {
				res, changes = wash(res)
				if changes == 0 {
					m[alpha] = len(res)
					break
				}
			}
		}(r)
	}

	// Wait for all goroutines to finish
	wg.Wait()

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
