package sumofparts

// Order returns the execution order of steps
func Order(steps []*Step) string {
	candidates := NewSteps(steps)
	for _, step := range steps {
		// Adds (root) steps which do not need any pre steps
		if len(step.preSteps) == 0 {
			candidates.Add(step)
		}
	}
	root, _ := candidates.Next()
	return traverse(root, candidates)
}

func traverse(current *Step, candidates Steps) string {
	current.done = true
	candidates.Remove(current)

	for _, step := range current.postSteps {
		// Add steps that are ready to be executed
		if step.Ready() {
			candidates.Add(step)
		}
	}

	next, err := candidates.Next()
	if err != nil {
		// We're out of candidate steps. Unfold recursion
		// by returning the key of the current step
		return current.key
	}
	return current.key + traverse(next, candidates)
}
