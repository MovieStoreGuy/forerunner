package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("We have an issue here: ", err)
		os.Exit(1)
	}
	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: "hello-world",
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		fmt.Println("Failed to run conainer due to:", err)
		os.Exit(-1)
	}
	if err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println("Failed to do the thing I love", err)
		os.Exit(-1)
	}
	// Perform actions that are required once the image is running
	t := 10 * time.Second
	if err = cli.ContainerStop(context.Background(), resp.ID, &t); err != nil {
		fmt.Println("Can not do it man:", err)
		os.Exit(-1)
	}
	if err = cli.ContainerRemove(context.Background(), resp.ID, types.ContainerRemoveOptions{Force: true}); err != nil {
		fmt.Println("I can not do that :", err)
		os.Exit(-1)
	}
}
