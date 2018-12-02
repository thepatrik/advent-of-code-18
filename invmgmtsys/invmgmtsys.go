package invmgmtsys

// ChkSum : Get the total checksum
func ChkSum(s []string) int {
	var twos, threes int
	for _, s := range s {
		tw, th := count(s)
		if tw > 0 {
			twos++
		}
		if th > 0 {
			threes++
		}
	}
	return twos * threes
}

// FindCmnLetters : Finds common IDs in string slice
func FindCmnLetters(ids []string) string {
	for i := range ids {
		for j := i + 1; j < len(ids); {
			common := comp(ids[i], ids[j])
			if len(common) > 0 {
				return common // return first hit
			}
			j++
		}
	}
	return ""
}

func comp(s1, s2 string) string {
	var common string
	for i := range s1 {
		if s1[i] == s2[i] {
			common = common + string(s1[i])
		}
	}

	if len(s1)-len(common) <= 1 {
		return common // common enough: 0-1 diff
	}
	return ""
}

func count(s string) (int, int) {
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
}
