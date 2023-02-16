package utils

import (
	"fmt"
	"log"
	"os"
)

// Logger ...
type Logger struct {
	k *log.Logger
}

// NewLogger ...
func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)}
}

func (c *Logger) print(s string) {
	_ = c.k.Output(3, s)
}

// Info ...
func (c *Logger) Info(s string, v ...any) {
	c.print("INFO: " + fmt.Sprintf(s, v...))
}

// Warn ...
func (c *Logger) Warn(s string, v ...any) {
	c.print("WARN: " + fmt.Sprintf(s, v...))
}

// Error ...
func (c *Logger) Error(s string, v ...any) {
	c.print("ERROR: " + fmt.Sprintf(s, v...))
}

// Fatal ...
func (c *Logger) Fatal(s string, v ...any) {
	c.print("FATAL: " + fmt.Sprintf(s, v...))
	os.Exit(1)
}
