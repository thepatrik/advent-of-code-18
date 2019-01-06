package sumofparts

// Order returns the execution order of steps
func Order(allsteps []*Step) string {
	steps := NewSteps()
	for _, step := range allsteps {
		// Add root steps (ie steps that do not need any pre steps)
		if step.IsRoot() {
			steps.Add(step)
		}
	}
	return iterate(steps)
}

func iterate(steps Steps) string {
	step, err := steps.Next()
	if err != nil {
		return "" // No more steps found
	}

	step.Done()
	steps.Remove(step)

	for _, post := range step.postSteps {
		// Add steps that are ready to be executed
		if post.IsReady() {
			steps.Add(post)
		}
	}
	return step.key + iterate(steps)
}
