package dullback

import (
	"log"
	"os"
)

type Stuff struct {
	logger Logger
}

func NewStuff(logger Logger) *Stuff {
	return &Stuff{logger: logger}
}

func (s *Stuff) DoStuff() {
	s.logger.Debug("Hello playground")
	s.logger.Debugf("Hello with %s", "formatting")
}

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type BuiltinLogger struct {
	logger *log.Logger
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "", 5)}
}

func (l *BuiltinLogger) Debug(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Debugf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}
