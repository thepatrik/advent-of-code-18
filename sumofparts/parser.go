package sumofparts

import "fmt"

func steps(strslice []string) []*Step {
	m := make(map[string]*Step)
	for _, s := range strslice {
		var stepstr string
		var prestepstr string
		_, _ = fmt.Sscanf(s, "Step %s must be finished before step %s can begin.",
			&prestepstr, &stepstr)

		preStep, ok := m[prestepstr]
		if !ok {
			preStep = NewStep(prestepstr)
			m[prestepstr] = preStep
		}

		step, ok := m[stepstr]
		if !ok {
			step = NewStep(stepstr)
			m[stepstr] = step
		}
		step.AddPreStep(preStep)
		preStep.AddPostStep(step)
	}
	steps := make([]*Step, 0, len(m))
	for _, v := range m {
		steps = append(steps, v)
	}
	return steps
}
