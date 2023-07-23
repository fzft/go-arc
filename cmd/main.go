package main

import (
	"github.com/fzft/go-arc/cli"
	"os"
)

var version string // set by linker, do not remove or modify

func main() {
	os.Exit(cli.Run())
}
