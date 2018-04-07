package runner

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"

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
	mode := container.NetworkMode(s.conf.Network)
	switch {
	case mode.IsBridge():
		fmt.Println("Is a Bridge setting")
	case mode.IsHost():
		fmt.Println("Is a host setting")
	default:
		fmt.Println("No fucking idea")
	}
	resp, err := s.cli.ContainerCreate(s.ctx, &container.Config{
		Image: image,
		Env:   s.conf.Environment,
	}, &container.HostConfig{NetworkMode: mode}, nil, "")
	if err != nil {
		return err
	}
	fmt.Println("Do I make it this far?")
	if err = s.cli.ContainerStart(s.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	s.ids = append(s.ids, resp.ID)
	return s.runCommands(s.conf.Cmds...)
}

func (s *Spartan) runCommands(cmds ...string) error {
	for _, str := range cmds {
		cmd := strings.Split(str, " ")
		c := exec.Command(cmd[0], cmd[1:]...)
		stdout, err := c.StdoutPipe()
		if err != nil {
			return errors.New("Unable to connect stdout")
		}
		stderr, err := c.StderrPipe()
		if err != nil {
			return errors.New("Unable to connect stderr")
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
		logs, err := s.cli.ContainerLogs(s.ctx, id, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
		if err != nil {
			return err
		}
		fmt.Println("Container logs:", id)
		writer(logs)
		if err := s.cli.ContainerStop(s.ctx, id, nil); err != nil {
			return err
		}
		if err := s.cli.ContainerRemove(s.ctx, id, types.ContainerRemoveOptions{Force: true}); err != nil {
			return err
		}
	}
	return nil
}

func writer(o io.ReadCloser) {
	scanner := bufio.NewScanner(o)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
