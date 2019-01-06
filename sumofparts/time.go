package sumofparts

const (
	alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// ToTime converts string to time
func ToTime(s string) int {
	r1 := []rune(s)[0]
	for i, r := range alpha {
		if r1 == r {
			return i + 1
		}
	}
	return -1
}
