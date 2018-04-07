package main

import (
	"fmt"
	"os"

	"github.com/MovieStoreGuy/forerunner/runner"
)

func main() {
	mchief, err := runner.New(nil)
	if err != nil {
		fmt.Println("An issued occured\n\t", err)
		os.Exit(1)
	}
	defer func() {
		if err := mchief.Stop(); err != nil {
			fmt.Println("Failed to stop runner\n\t", err)
			os.Exit(-1)
		}
	}()
	for _, image := range os.Args[1:] {
		if err = mchief.Start(image); err != nil {
			fmt.Println("Failed to start the runner\n\t", err)
			return
		}
	}
}
