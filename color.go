package goarc

import (
	"github.com/fatih/color"
	"io"
)

type StdWriter interface {
	Printf(format string, a ...any)
	Println(a ...any)
}

// StdOutWriter is a StdWriter that writes to Stdout
type StdOutWriter struct {
	out   io.Writer
	color *color.Color
}

func NewStdOutWriter(out io.Writer, color *color.Color) StdWriter {
	return &StdOutWriter{out: out, color: color}
}

func DefaultStdOutWriter(out io.Writer) StdWriter {
	return NewStdOutWriter(out, color.New(color.FgCyan))
}

func (s *StdOutWriter) Printf(format string, a ...any) {
	s.color.Fprint(s.out, format, a)
}

func (s *StdOutWriter) Println(a ...any) {
	s.color.Fprintln(s.out, a)
}

// StdErrWriter is a StdWriter that writes to Stderr
type StdErrWriter struct {
	err   io.Writer
	color *color.Color
}

func NewStdErrWriter(err io.Writer, color *color.Color) StdWriter {
	return &StdErrWriter{err: err, color: color}
}

func DefaultStdErrWriter(err io.Writer) StdWriter {
	return NewStdErrWriter(err, color.New(color.FgRed))
}

func (s *StdErrWriter) Printf(format string, a ...any) {
	s.color.Fprint(s.err, format, a)
}

func (s *StdErrWriter) Println(a ...any) {
	s.color.Fprintln(s.err, a)
}
