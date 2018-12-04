package nomatterhowyousliceit

// GetOverlappingFabrics : Get the number of overlapping fabrics
func GetOverlappingFabrics(slice []string) int {
	fabric := GetFabric(slice)

	count := 0
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric); j++ {
			if fabric[j][i] > 1 {
				count++
			}
		}
	}

	return count
}
