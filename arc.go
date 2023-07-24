package goarc

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

type Arc struct {
	outWriter, errWriter StdWriter
}

func New(out, err io.Writer) *Arc {
	return &Arc{outWriter: DefaultStdOutWriter(out), errWriter: DefaultStdErrWriter(err)}
}

// Run starts the application,
func (a *Arc) Run(_ context.Context) error {
	a.outWriter.Println("hello, welcome to go-arc!")

	// interact with user, collect answers
	for _, qa := range *QAs {
		if err := qa.Ask(); err != nil && strings.Contains(err.Error(), "interrupt") {
			os.Exit(1)
		}
	}

	for _, qa := range *QAs {
		fmt.Printf("Answer: %v\n", qa.Answer)
	}

	return nil
}
