package invmgmtsys

// ChkSum : Get the total checksum
func ChkSum(s []string) int {
	var twos, threes int
	for _, s := range s {
		tw, th := func(s string) (int, int) {
			m := make(map[string]int)
			for _, rune := range s {
				m[string(rune)]++
			}

			var twos, threes int
			for _, v := range m {
				switch v {
				case 2:
					twos++
				case 3:
					threes++
				}
			}
			return twos, threes
		}(s)
		if tw > 0 {
			twos++
		}
		if th > 0 {
			threes++
		}
	}
	return twos * threes
}
