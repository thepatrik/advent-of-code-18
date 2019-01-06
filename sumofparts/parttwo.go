package sumofparts

// CompletionTime returns time it takes to complete all steps
func CompletionTime(allsteps []*Step, concurrency int) int {
	steps := NewSteps()
	for _, step := range allsteps {
		// Adds (root) steps which do not need any pre steps
		if len(step.preSteps) == 0 {
			steps.Add(step)
		}
	}
	worker := NewWorker(concurrency)
	return processTime(steps, worker)
}

func processTime(steps Steps, worker *Worker) int {
	if steps.Len() == 0 {
		return 0 // No more steps to process
	}

	all := steps.All()
	worker.Deploy(all)

	// Let the worker process steps and capture time
	time, processed := worker.Process()

	// Removes processed steps
	steps.RemoveSteps(processed)

	// Add new steps ready for processing
	for _, step := range all {
		for _, post := range step.postSteps {
			if post.IsReady() {
				steps.Add(post)
			}
		}
	}

	return time + processTime(steps, worker)
}
