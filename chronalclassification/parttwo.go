package chronalclassification

const noOfOperations = 16

// CalculateRegistry calculates the value of the last registry
func CalculateRegistry(samples []*Sample, instructions []Instruction) int {
	ops := operations(samples)
	reg := []int{0, 0, 0, 0}

	for _, instr := range instructions {
		code, _ := (*ops)[instr.Code]
		sample := Sample{}
		sample.RegBefore = reg
		sample.Instr = instr
		switch code {
		case "addr":
			reg = sample.addr()
		case "addi":
			reg = sample.addi()
		case "mulr":
			reg = sample.mulr()
		case "muli":
			reg = sample.muli()
		case "banr":
			reg = sample.banr()
		case "bani":
			reg = sample.bani()
		case "borr":
			reg = sample.borr()
		case "bori":
			reg = sample.bori()
		case "setr":
			reg = sample.setr()
		case "seti":
			reg = sample.seti()
		case "gtir":
			reg = sample.gtir()
		case "gtri":
			reg = sample.gtri()
		case "gtrr":
			reg = sample.gtrr()
		case "eqir":
			reg = sample.eqir()
		case "eqri":
			reg = sample.eqri()
		case "eqrr":
			reg = sample.eqrr()
		}
	}
	return reg[0]
}

func operations(samples []*Sample) *map[int]string {
	ops := make(map[int]map[string]int)
	for _, s := range samples {
		funcs := make(map[string]operation)
		funcs["addr"] = s.addr
		funcs["addi"] = s.addi
		funcs["mulr"] = s.mulr
		funcs["muli"] = s.muli
		funcs["banr"] = s.banr
		funcs["bani"] = s.bani
		funcs["borr"] = s.borr
		funcs["bori"] = s.bori
		funcs["setr"] = s.setr
		funcs["seti"] = s.seti
		funcs["gtir"] = s.gtir
		funcs["gtri"] = s.gtri
		funcs["gtrr"] = s.gtrr
		funcs["eqir"] = s.eqir
		funcs["eqri"] = s.eqri
		funcs["eqrr"] = s.eqrr

		for k, f := range funcs {
			if s.eq(f()) {
				_, ok := ops[s.Instr.Code]
				if !ok {
					ops[s.Instr.Code] = make(map[string]int)
				}
				ops[s.Instr.Code][k]++
			}
		}
	}

	res := make(map[int]string)
	for len(res) < noOfOperations {
		k, v := func() (int, string) {
			for k, v := range ops {
				if len(v) == 1 {
					for k2 := range v {
						return k, k2
					}
				}
			}
			return 0, ""
		}()
		res[k] = v

		for k2, v2 := range ops {
			for k3 := range v2 {
				if k3 == v {
					delete(ops[k2], k3)
				}
			}
		}
	}

	return &res
}
