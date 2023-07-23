package cli

import (
	"flag"
	"github.com/fzft/go-arc"
	"io"
	"os"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErr
)

func Run() int {
	return (&cli{outWriter: os.Stdout, errWriter: os.Stderr}).run(os.Args[1:])
}

type cli struct {
	outWriter, errWriter io.Writer
}

func (c *cli) parseArgs(args []string) (*goarc.Arc, error) {
	fs := flag.NewFlagSet("go-arc", flag.ExitOnError)
	fs.SetOutput(c.errWriter)
	fs.Usage = func() {
		fs.SetOutput(c.outWriter)
		defer fs.SetOutput(c.errWriter)
		fs.PrintDefaults()
	}
	return goarc.New(c.outWriter, c.errWriter), nil
}

func (c *cli) run(args []string) int {
	arc, err := c.parseArgs(args)
	if err != nil {
		return int(exitCodeErr)
	}

	if err = arc.Run(); err != nil {
		return int(exitCodeErr)
	}

	return int(exitCodeOK)
}
