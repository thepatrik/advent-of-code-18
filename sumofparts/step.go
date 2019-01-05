package sumofparts

// Step type
type Step struct {
	key       string
	done      bool
	preSteps  []*Step
	postSteps []*Step
}

// NewStep constructs a new step
func NewStep(key string) *Step {
	preSteps := make([]*Step, 0)
	postSteps := make([]*Step, 0)
	return &Step{key: key, preSteps: preSteps, postSteps: postSteps}
}

// AddPreStep will add a step as a pre step
func (step *Step) AddPreStep(preSteps *Step) {
	step.preSteps = append(step.preSteps, preSteps)
}

// AddPostStep will add a step as a post step
func (step *Step) AddPostStep(postSteps *Step) {
	step.postSteps = append(step.postSteps, postSteps)
}

// Ready return true if all pre steps are done
func (step *Step) Ready() bool {
	for _, step := range step.preSteps {
		if !step.done {
			return false
		}
	}
	return true
}
