package runner

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/MovieStoreGuy/forerunner/config"
	"github.com/MovieStoreGuy/forerunner/cortana"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Spartan is the object type to allow management of
// running the docker image
type Spartan struct {
	cli     *client.Client
	ctx     context.Context
	conf    *config.Set
	ids     []string
	Cortana *cortana.Sentient
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
		cli:     cli,
		ctx:     context.Background(),
		conf:    conf,
		Cortana: cortana.New(),
	}, nil
}

// Start will run the docker image
func (s *Spartan) Start(image string) error {
	if !s.containerExists(image) {
		return fmt.Errorf("The container %s doesn't appear to exist", image)
	}
	mode := container.NetworkMode(s.conf.Network)
	resp, err := s.cli.ContainerCreate(s.ctx, &container.Config{
		Image: image,
		Env:   s.conf.Environment,
	}, &container.HostConfig{NetworkMode: mode}, nil, "")
	if err != nil {
		return err
	}
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
		go s.Cortana.Follow(stdout)
		go s.Cortana.Follow(stderr)
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
		s.Cortana.Follow(logs)
		if err := s.cli.ContainerStop(s.ctx, id, nil); err != nil {
			return err
		}
		if err := s.cli.ContainerRemove(s.ctx, id, types.ContainerRemoveOptions{Force: true}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Spartan) containerExists(image string) bool {
	images, err := s.cli.ImageList(s.ctx, types.ImageListOptions{All: true})
	if err != nil {
		return false
	}
	for _, sum := range images {
		if strings.Contains(strings.Join(sum.RepoTags, " "), image) {
			return true
		}
	}
	return false
}
