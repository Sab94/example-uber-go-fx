package loggerfx

import (
	"log"
	"os"

	"go.uber.org/fx"
)

var Module = fx.Options(
	// Provide constructors for loggerfx
	fx.Provide(
		NewLogger,
	),
)

// Logger Interface
type Logger interface {
	Println(v ...interface{})
}

// Logger constructor
func NewLogger() Logger {
	return log.New(os.Stdout, "FX ", 0)
}
