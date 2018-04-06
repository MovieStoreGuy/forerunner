package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("We have an issue here: ", err)
		os.Exit(1)
	}
	resp, err := cli.ContainerExecCreate(context.Background(), "hello-world", types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
	})
	if err != nil {
		fmt.Println("Today is a day of sadness:", err)
		os.Exit(1)
	}
	_ = resp
}
