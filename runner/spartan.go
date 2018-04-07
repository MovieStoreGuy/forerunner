package runner

// Spartan is the object type to allow management of
// running the docker image
type Spartan struct {
}

// New will create an instance of Spartan with the desired config
func New(config ...interface{}) *Spartan {
	runner := &Spartan{}
	// Do some magic with config
	return runner
}

// Start will run the docker image
func (s *Spartan) Start(image string) error {

	return nil
}

func (s *Spartan) runCommands(cmds ...[]string) error {
	for _, cmd := range cmds {
		_ = cmd
	}
	return nil
}

// Stop will stop the given image from running and ensure that anything else
// created with it is also cleaned up
func (s *Spartan) Stop(image string) error {
	return nil
}
