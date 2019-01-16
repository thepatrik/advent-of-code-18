package dartminemadness

// Direction type
type Direction int

const (
	north Direction = iota
	south
	west
	east
	none
)

// Turn gets the new direction after a turn
func (direction Direction) Turn(t turn) Direction {
	dir := direction
	switch direction {
	case west:
		switch t {
		case right:
			dir = north
		case left:
			dir = south
		}
	case east:
		switch t {
		case right:
			dir = south
		case left:
			dir = north
		}
	case south:
		switch t {
		case right:
			dir = west
		case left:
			dir = east
		}
	case north:
		switch t {
		case right:
			dir = east
		case left:
			dir = west
		}
	}

	return dir
}

func (direction Direction) String() string {
	switch direction {
	case north:
		return "^"
	case south:
		return "v"
	case west:
		return "<"
	case east:
		return ">"
	}
	return ""
}
