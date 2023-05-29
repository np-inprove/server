package logger

import (
	"github.com/charmbracelet/log"
	"os"
)

type charm struct {
	l *log.Logger
}

func NewCharm() AppLogger {
	l := log.New(os.Stderr)
	return charm{l}
}

func expandZapFields(f []Field) []interface{} {
	fields := make([]interface{}, len(f)*2)
	for i, field := range f {
		fields[2*i] = field.Key
		if field.String == "" {
			fields[(2*i)+1] = field.Integer
		} else {
			fields[(2*i)+1] = field.String
		}
	}
	return fields
}

func (c charm) Debug(msg string, fields ...Field) {
	c.l.Debug(msg, expandZapFields(fields)...)
}

func (c charm) Info(msg string, fields ...Field) {
	c.l.Info(msg, expandZapFields(fields)...)
}

func (c charm) Warn(msg string, fields ...Field) {
	c.l.Warn(msg, expandZapFields(fields)...)
}

func (c charm) Error(msg string, fields ...Field) {
	c.l.Error(msg, expandZapFields(fields)...)
}

func (c charm) DPanic(msg string, fields ...Field) {
	c.l.Fatal(msg, expandZapFields(fields)...)
}

func (c charm) Panic(msg string, fields ...Field) {
	c.l.Fatal(msg, expandZapFields(fields)...)
}

func (c charm) Fatal(msg string, fields ...Field) {
	c.l.Fatal(msg, expandZapFields(fields)...)
}
