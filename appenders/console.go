package appenders

import (
	"fmt"

	"github.com/lakshay2395/go-log/layout"
	"github.com/lakshay2395/go-log/levels"
)

type consoleAppender struct {
	Appender
	layout layout.Layout
}

func Console() *consoleAppender {
	a := &consoleAppender{
		layout: layout.Default(),
	}
	return a
}

func (a *consoleAppender) Write(level levels.LogLevel, message string, args ...interface{}) {
	fmt.Println(a.Layout().Format(level, message, args...))
}

func (a *consoleAppender) Layout() layout.Layout {
	return a.layout
}

func (a *consoleAppender) SetLayout(layout layout.Layout) {
	a.layout = layout
}
