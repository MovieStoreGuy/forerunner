package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MovieStoreGuy/forerunner/config"
	"github.com/MovieStoreGuy/forerunner/runner"
)

var (
	conf = config.Default()

	confpath string
)

func init() {
	flag.StringVar(&confpath, "path", "", "allows the user to define their own runner config instead of the default")
}

func main() {
	flag.Parse()
	if confpath != "" {
		c, err := config.Load(confpath)
		if err != nil {
			fmt.Println("Unable to load config\n\t", err)
			os.Exit(-1)
		}
		conf = c
	}
	fmt.Printf("Config is: %+v\n", conf)
	mchief, err := runner.New(conf)
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
