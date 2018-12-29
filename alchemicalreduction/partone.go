package alchemicalreduction

// CalculateUnits calculates how many units
// remain after fully reacting the polymer.
func CalculateUnits(s string) int {
	res := []rune(s)
	changes := 0
	for {
		res, changes = wash(res)
		if changes == 0 {
			return len(res)
		}
	}
}
