package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
)


func main() {
	app := &cli.App{
		Action: func(cCtx *cli.context) error {
			fmt.Println("Hello %q", cCtx.Args().Get(0))
			return nil
		},
	}
	
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

