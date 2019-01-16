package dartminemadness

// LastCartPos returns the position for the last surviving cart.
func LastCartPos(raceTrack *RaceTrack) Position {
	if raceTrack.carts.Len() == 1 {
		// Only one cart left on the race track, return it's
		// position.
		return *raceTrack.carts.Get(0).pos
	}
	if hit, wrecks := raceTrack.Wrecks(); hit {
		// Get all wrecks and remove them before next round.
		raceTrack.carts.RemoveAll(wrecks)
	}
	return LastCartPos(raceTrack)
}
