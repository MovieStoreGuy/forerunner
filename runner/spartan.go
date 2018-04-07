package runner

import (
	"context"

	"github.com/MovieStoreGuy/forerunner/config"
	"github.com/docker/docker/client"
)

// Spartan is the object type to allow management of
// running the docker image
type Spartan struct {
	cli  *client.Client
	ctx  context.Context
	conf *config.Set
}

// New will create an instance of Spartan with the desired config
func New(conf *config.Set) (*Spartan, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	if conf == nil {
		conf = config.Default()
	}
	return &Spartan{
		cli:  cli,
		ctx:  context.Background(),
		conf: conf,
	}, nil
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
