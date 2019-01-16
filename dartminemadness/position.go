package dartminemadness

import "strconv"

// Position (x, y) type
type Position struct {
	x int
	y int
}

// String returns a string representation
func (pos Position) String() string {
	return strconv.Itoa(pos.x) + ":" + strconv.Itoa(pos.y)
}
