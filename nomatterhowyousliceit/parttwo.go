package nomatterhowyousliceit

// GetNonOverlappingFabricID : Get the IDs of fabrics that are not overlapping any other fabric
func GetNonOverlappingFabricID(slice []string) int {
	fabric := GetFabric(slice)

	for _, s := range slice {
		n, x, y, w, h := GetValues(s)

		if found := func() bool {
			for dx := 0; dx < w; dx++ {
				for dy := 0; dy < h; dy++ {
					if fabric[x+dx][y+dy] > 1 {
						return false
					}
				}
			}
			return true
		}(); found {
			return n
		}
	}

	return -1
}
