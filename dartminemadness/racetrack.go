package dartminemadness

// RaceTrack type
type RaceTrack struct {
	tracks *Tracks
	carts  *Carts
}

// NewRaceTrack constructs a new race track
func NewRaceTrack(tracks *Tracks, carts *Carts) *RaceTrack {
	return &RaceTrack{tracks: tracks, carts: carts}
}

// Tick ticks the race track one round and if two carts have crashed
// will return the position they crashed on
func (raceTrack *RaceTrack) Tick() (bool, *Position) {
	sorted := raceTrack.carts.Sort()
	for _, cart := range sorted {
		raceTrack.Move(cart)
		if hit, _ := raceTrack.carts.HitTest(cart); hit {
			return true, cart.pos
		}
	}
	return false, nil
}

// Wrecks finds all the wrecks (crashed cars) in a tick round
func (raceTrack *RaceTrack) Wrecks() (bool, []*Cart) {
	totalwrecks := make([]*Cart, 0)

	sorted := raceTrack.carts.Sort()
	for _, cart := range sorted {
		raceTrack.Move(cart)
		if hit, wrecks := raceTrack.carts.HitTest(cart); hit {
			for _, wreck := range wrecks {
				totalwrecks = append(totalwrecks, wreck)
			}
		}
	}

	res := false
	if len(totalwrecks) > 0 {
		res = true
	}
	return res, totalwrecks
}

// Move moves a cart on the track
func (raceTrack *RaceTrack) Move(cart *Cart) {
	switch raceTrack.tracks.Type(*cart.pos) {
	case slash:
		switch cart.direction {
		case east: // >/
			cart.Move(north)
		case south: // v/
			cart.Move(west)
		case north: // /^
			cart.Move(east)
		case west: // /<
			cart.Move(south)
		}
	case backslash:
		switch cart.direction {
		case east: // >\
			cart.Move(south)
		case south: // \v
			cart.Move(east)
		case west: // \<
			cart.Move(north)
		case north: // ^\
			cart.Move(west)
		}
	case intersection:
		cart.direction = cart.Intersection()
		fallthrough
	default:
		cart.Move(cart.direction)
	}
}

// String converts tracks and carts to plottable string
func (raceTrack RaceTrack) String() string {
	str := "\n"
	slice := raceTrack.tracks.Slice()
	for y, s := range slice {
		for x, r := range s {
			cart, err := raceTrack.carts.FindCart(x, y)
			if err != nil {
				str += string(r)
			} else {
				str += cart.direction.String()
			}
		}
		str += "\n"
	}
	return str
}
