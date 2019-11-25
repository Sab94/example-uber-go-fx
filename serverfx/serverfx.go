package serverfx

import (
	"context"
	"net/http"

	"github.com/eg-fx/loggerfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// Constructers are called lazity. So Invoke the Handler and Init Server to kick start the app
	fx.Invoke(
		RegisterHandlers,
		InitServer,
	),
)

func RegisterHandlers(handler http.Handler) {
	http.Handle("/", handler)
}

func InitServer(lc fx.Lifecycle, logger loggerfx.Logger) {
	server := &http.Server{
		Addr: ":8080",
	}

	// fx.Hook has two callbacks OnStart and OnStop. They do what their name implies
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Println("OnStart Called")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Println("OnStop Called")
			return server.Close()
		},
	})
}
