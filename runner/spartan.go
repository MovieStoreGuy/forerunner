package runner

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"

	"github.com/MovieStoreGuy/forerunner/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Spartan is the object type to allow management of
// running the docker image
type Spartan struct {
	cli  *client.Client
	ctx  context.Context
	conf *config.Set
	ids  []string
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
	resp, err := s.cli.ContainerCreate(s.ctx, &container.Config{
		Image: image,
	}, nil, nil, "")
	if err != nil {
		return err
	}
	if err = s.cli.ContainerStart(s.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	s.ids = append(s.ids, resp.ID)
	return s.runCommands(s.conf.Cmds...)
}

func (s *Spartan) runCommands(cmds ...[]string) error {
	for _, cmd := range cmds {
		c := exec.Command(cmd[0], cmd[1:]...)
		stdout, err := c.StdoutPipe()
		if err != nil {
			return errors.New("Unable to connect stdout")
		}
		stderr, err := c.StderrPipe()
		if err != nil {
			return errors.New("Unable to connect stderr")
		}
		writer := func(o io.ReadCloser) {
			scanner := bufio.NewScanner(o)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
		// Write outputs as soon as they are received
		go writer(stdout)
		go writer(stderr)
		if err = c.Run(); err != nil {
			return err
		}
	}
	return nil
}

// Stop will stop the given image from running and ensure that anything else
// created with it is also cleaned up
func (s *Spartan) Stop() error {
	for _, id := range s.ids {
		if err := s.cli.ContainerStop(s.ctx, id, nil); err != nil {
			return err
		}
		if err := s.cli.ContainerRemove(s.ctx, id, types.ContainerRemoveOptions{Force: true}); err != nil {
			return err
		}
	}
	return nil
}
