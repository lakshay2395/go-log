package layout

import (
	"fmt"

	"github.com/lakshay2395/go-log/levels"
)

type basicLayout struct {
	Layout
}

func Basic() *basicLayout {
	return &basicLayout{}
}

func (a *basicLayout) Format(level levels.LogLevel, message string, args ...interface{}) string {
	return fmt.Sprintf(message, args...)
}
