package dartminemadness

// Type type
type Type int

const (
	empty Type = iota
	vertical
	horizontal
	slash
	backslash
	intersection
)

// String returns a string representation
func (t Type) String() string {
	s := ""
	switch t {
	case empty:
		s = " "
	case vertical:
		s = "|"
	case horizontal:
		s = "-"
	case slash:
		s = "/"
	case backslash:
		s = "\\"
	case intersection:
		s = "+"
	}
	return s
}
