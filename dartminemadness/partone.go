package dartminemadness

// DetectFirstCrash finds position of the first crash
func DetectFirstCrash(raceTrack *RaceTrack) *Position {
	if hit, pos := raceTrack.Tick(); hit {
		// We have found a first hit. Return that position.
		return pos
	}
	return DetectFirstCrash(raceTrack)
}
