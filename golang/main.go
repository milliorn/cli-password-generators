package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func main() {
	// fmt.Println("Hello, World!")
	app := &cli.App{
		Name:  "hello",
		Usage: "Print Hello World",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello, World!")
			return nil
		},
	}
	app.Run([]string{"hello"})
}
