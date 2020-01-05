package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		var cmd string
		fmt.Printf("%T\n", c.Args())
		if c.NArg() > 0 {
			//cmd = c.Args()[0]
		}
		fmt.Println("hello friend", cmd)
		return nil
	}
	app.Run(os.Args)
}
