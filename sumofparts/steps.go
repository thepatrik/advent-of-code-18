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
func NewSteps() Steps {
	m := make(map[string]*Step)
	return Steps{m: m}
}

// Len returns the number of steps
func (steps *Steps) Len() int {
	return len(steps.m)
}

// Add adds a step
func (steps *Steps) Add(step *Step) {
	steps.m[step.key] = step
}

// RemoveSteps removes slice of steps
func (steps *Steps) RemoveSteps(stepslice []*Step) {
	for _, step := range stepslice {
		delete(steps.m, step.key)
	}
}

// Remove removes a step
func (steps *Steps) Remove(step *Step) {
	delete(steps.m, step.key)
}

// All get all steps
func (steps *Steps) All() []*Step {
	res := make([]*Step, 0, len(steps.m))
	for _, v := range steps.m {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].time < res[j].time
	})

	return res
}

// Next finds the next step. If there are multiple steps,
// the next step will be the first in alpha order of the key
func (steps *Steps) Next() (*Step, error) {
	if steps.Len() == 0 {
		return &Step{}, errors.New("Nothing to be done")
	}

	return steps.All()[0], nil
}
