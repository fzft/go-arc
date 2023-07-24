package cli

import (
	"context"
	"flag"
	"fmt"
	"github.com/fzft/go-arc"
	"io"
	"os"
	"os/signal"
	"syscall"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErr
)

func Run() int {
	return (&cli{outWriter: os.Stdout, errWriter: os.Stderr, sigChan: make(chan os.Signal)}).run(os.Args[1:])
}

type cli struct {
	outWriter, errWriter io.Writer
	sigChan              chan os.Signal
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration files
	goarc.LoadConf()

	// Load questions
	goarc.QAs.LoadQAs()

	// Handle signals
	signal.Notify(c.sigChan, syscall.SIGINT, syscall.SIGTERM)

	go arc.Run(ctx)

	// Wait for signals
	sig := <-c.sigChan
	fmt.Printf("Received signal: %s. Exiting...\n", sig)

	return int(exitCodeOK)
}
