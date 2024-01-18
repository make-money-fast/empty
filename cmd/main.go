package main

import (
	"context"
	"github.com/make-money-fast/empty/config"
	"github.com/make-money-fast/empty/internal/model"
	"github.com/make-money-fast/empty/internal/router"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Supply(config.Load()),
		fx.Provide(model.Load),
		fx.Invoke(func(lifecycle fx.Lifecycle, conf *config.Config) {
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return router.Start(conf)
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),
	)
	app.Start(context.Background())
}
