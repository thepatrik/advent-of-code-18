package sumofparts

import (
	"errors"
	"sort"
)

// Steps type
type Steps struct {
	m map[string]*Step
}

// NewSteps constructs a new steps type
func NewSteps(steps []*Step) Steps {
	m := make(map[string]*Step)
	return Steps{m: m}
}

// Add adds a step
func (steps *Steps) Add(step *Step) {
	steps.m[step.key] = step
}

// Remove removes a step
func (steps *Steps) Remove(step *Step) {
	delete(steps.m, step.key)
}

// Next finds the next step. If there are multiple steps,
// the next step will be the first in alpha order of the key
func (steps *Steps) Next() (*Step, error) {
	keys := make([]string, 0, len(steps.m))
	for k := range steps.m {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return &Step{}, errors.New("Nothing to be done")
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return steps.m[keys[0]], nil
}
