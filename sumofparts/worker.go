package sumofparts

import (
	"errors"
	"sort"
)

// Worker type
type Worker struct {
	steps []*Step
}

// NewWorker constructs a new worker
func NewWorker(concurrency int) *Worker {
	steps := make([]*Step, concurrency)
	return &Worker{steps: steps}
}

// Deploy adds steps to worker
func (worker *Worker) Deploy(steps []*Step) {
	for _, step := range steps {
		if !worker.IsDeployed(step) {
			i, err := worker.FindSpot()
			if err == nil {
				worker.steps[i] = step
			}
		}
	}
}

// Clean cleans up finished steps
func (worker *Worker) Clean() []*Step {
	cleaned := make([]*Step, 0)
	for i, step := range worker.steps {
		if step != nil && step.IsDone() {
			cleaned = append(cleaned, step)
			worker.steps[i] = nil
		}
	}
	return cleaned
}

// FindSpot finds an available worker spot
func (worker *Worker) FindSpot() (int, error) {
	for i, step := range worker.steps {
		if step == nil {
			return i, nil
		}
	}
	return -1, errors.New("No spot found")
}

// IsDeployed returns true if step is deployed
func (worker *Worker) IsDeployed(step *Step) bool {
	for _, wstep := range worker.steps {
		if wstep != nil && wstep.key == step.key {
			return true
		}
	}
	return false
}

// Sort sorts steps in time order
func (worker *Worker) Sort() {
	sort.Slice(worker.steps, func(i, j int) bool {
		if worker.steps[i] == nil {
			return false
		}
		if worker.steps[j] == nil {
			return true
		}
		return worker.steps[i].time < worker.steps[j].time
	})
}

// Size returns number of occupied workers
func (worker *Worker) Size() int {
	size := 0
	for _, step := range worker.steps {
		if step != nil {
			size++
		}
	}
	return size
}

// Process returns processing time and proccesses steps
func (worker *Worker) Process() (int, []*Step) {
	if worker.Size() == 0 {
		return 0, []*Step{}
	}

	// Sort steps after time asc
	worker.Sort()

	// The first step will be the shortest one
	time := worker.steps[0].Done()

	// Tick the rest with the same time
	for i := 1; i < len(worker.steps); i++ {
		step := worker.steps[i]
		if step != nil {
			step.Tick(time)
		}
	}

	return time, worker.Clean()
}
