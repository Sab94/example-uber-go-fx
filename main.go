package main

import (
	"github.com/eg-fx/handlers"
	"github.com/eg-fx/loggerfx"
	"github.com/eg-fx/serverfx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		serverfx.Module,
		loggerfx.Module,
		fx.Provide(
			handlers.NewHandler,
		),
	)

	app.Run()
}
