package chronalclassification

// CalculateSamples calculates how many samples behaves
// like three or more opt codes
func CalculateSamples(samples []*Sample) int {
	total := 0
	for _, s := range samples {
		c := 0

		funcs := []operation{s.addr, s.addi, s.mulr, s.muli,
			s.banr, s.bani, s.borr, s.bori,
			s.setr, s.seti, s.gtir, s.gtri,
			s.gtrr, s.eqir, s.eqri, s.eqrr}

		for _, f := range funcs {
			if s.eq(f()) {
				c++
			}
		}
		if c >= 3 {
			total++
		}
	}
	return total
}
