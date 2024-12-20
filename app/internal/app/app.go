package app

import (
	"context"
	"dzrise.ru/internal/pkg/closer"
	"fmt"
)

type App struct {
	sp *ServiceProvider
}

func New() (*App, error) {
	sp := NewServiceProvider()

	sp.Logger().Info("Initializing app...")
	sp.Logger().Info("Config", sp.Config())

	return &App{
		sp: sp,
	}, nil
}

// Run starts the HTTP server and listens for shutdown signals.
func (a *App) Run(ctx context.Context) error {
	a.sp.Logger().Info("Starting server...")

	defer func() {
		closer.CloseAll()
		closer.Wait()
		a.sp.Logger().Info("Server stopped.")
	}()

	addr := fmt.Sprintf("%s:%s", a.sp.Config().Server.Host, a.sp.Config().Server.Port)

	return a.sp.GetServer(ctx).Listen(addr)

}
