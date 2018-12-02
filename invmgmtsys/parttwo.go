package invmgmtsys

// FindCmnLetters : Finds common IDs in string slice
func FindCmnLetters(ids []string) string {
	for i := range ids {
		for j := i + 1; j < len(ids); {
			if comp := func(s1, s2 string) string {
				var curr string
				for i := range s1 {
					if s1[i] == s2[i] {
						curr = curr + string(s1[i])
					}
				}
				if len(s1)-len(curr) <= 1 {
					return curr // common enough: 0-1 diff
				}
				return ""
			}(ids[i], ids[j]); len(comp) > 0 {
				return comp // return first hit
			}
			j++
		}
	}
	return ""
}
