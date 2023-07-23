package goarc

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type Arc struct {
	outWriter, errWriter StdWriter
}

func New(out, err io.Writer) *Arc {
	return &Arc{outWriter: DefaultStdOutWriter(out), errWriter: DefaultStdErrWriter(err)}
}

// Run starts the application,
func (a *Arc) Run() error {
	scanner := bufio.NewScanner(os.Stdin)
	a.outWriter.Println("hello, welcome to go-arc!")

	for {
		a.outWriter.Println("Choose an option:")
		a.outWriter.Println("1. First choice")
		a.outWriter.Println("2. Second choice")
		a.outWriter.Println("3. Third choice")
		a.outWriter.Println("4. Exit")

		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			a.errWriter.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			a.outWriter.Println("First choice")
		case 2:
			a.outWriter.Println("Second choice")
		case 3:
			a.outWriter.Println("Third choice")
		case 4:
			a.outWriter.Println("Goodbye!")
			return nil
		default:
			a.errWriter.Println("Invalid choice")
		}
	}
}
