package dartminemadness

// Tracks type
type Tracks struct {
	types [][]Type
}

// NewTracks constructs new tracks type
func NewTracks(types [][]Type) *Tracks {
	return &Tracks{types: types}
}

// Type returns type at position
func (tracks *Tracks) Type(p Position) Type {
	return tracks.types[p.y][p.x]
}

// Slice returns string slice
func (tracks Tracks) Slice() []string {
	slice := make([]string, 0)
	for _, x := range tracks.types {
		s := ""
		for _, y := range x {
			s += y.String()
		}
		slice = append(slice, s)
	}
	return slice
}

// String returns a string representation
func (tracks Tracks) String() string {
	slice := tracks.Slice()
	res := "\n"
	for _, s := range slice {
		res += s + "\n"
	}
	return res
}
