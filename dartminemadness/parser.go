package dartminemadness

import "strconv"

func raceTrack(strslice []string) *RaceTrack {
	types := make([][]Type, len(strslice))
	carts := NewCarts()
	id := 0
	for y, s := range strslice {
		slice := make([]Type, 0, len(s))
		for x, r := range s {
			t, cartdirection := tracktypes(r)
			if cartdirection != none {
				cart := NewCart(cartdirection, strconv.Itoa(x)+":"+strconv.Itoa(y), x, y)
				carts.Add(cart)
				id++
			}
			slice = append(slice, t)
		}
		types[y] = slice
	}
	tracks := NewTracks(types)
	return NewRaceTrack(tracks, carts)
}

func tracktypes(r rune) (Type, Direction) {
	tracktype := empty
	direction := none
	switch r {
	case '|':
		tracktype = vertical
	case '-':
		tracktype = horizontal
	case '+':
		tracktype = intersection
	case '/':
		tracktype = slash
	case '\\':
		tracktype = backslash
	case '^':
		tracktype = vertical
		direction = north
	case 'v':
		tracktype = vertical
		direction = south
	case '>':
		tracktype = horizontal
		direction = east
	case '<':
		tracktype = horizontal
		direction = west
	}
	return tracktype, direction
}
