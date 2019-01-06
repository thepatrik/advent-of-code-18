package sumofparts

// Step type
type Step struct {
	key       string
	time      int
	preSteps  []*Step
	postSteps []*Step
}

// NewStep constructs a new step
func NewStep(key string) *Step {
	preSteps := make([]*Step, 0)
	postSteps := make([]*Step, 0)
	time := 60 + ToTime(key)
	return &Step{key: key, preSteps: preSteps, postSteps: postSteps, time: time}
}

// Tick ticks time
func (step *Step) Tick(time int) int {
	step.time -= time
	return step.time
}

// AddPreStep will add a step as a pre step
func (step *Step) AddPreStep(preSteps *Step) {
	step.preSteps = append(step.preSteps, preSteps)
}

// AddPostStep will add a step as a post step
func (step *Step) AddPostStep(postSteps *Step) {
	step.postSteps = append(step.postSteps, postSteps)
}

// IsRoot returns if step is a root step or not (ie requires
// no pre steps)
func (step *Step) IsRoot() bool {
	return len(step.preSteps) == 0
}

// Done will set the step as finished
func (step *Step) Done() int {
	delta := step.time
	step.time = 0
	return delta
}

// IsDone returns true if done
func (step *Step) IsDone() bool {
	return step.time == 0
}

// IsReady return true if all pre steps are done
func (step *Step) IsReady() bool {
	for _, step := range step.preSteps {
		if !step.IsDone() {
			return false
		}
	}
	return true
}
