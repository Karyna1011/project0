package main

import (
	"os"

  "gitlab.com/tokend/subgroup/project/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
