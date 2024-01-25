package protocol

import (
	"context"
	"database/sql"
	"time"

	"github.com/sifer169966/go-logger"

	"github.com/labstack/echo/v4"
)

type gracefulDependencies struct {
	srv      *echo.Echo
	postgres *sql.DB

	errCh chan error
}

type graceful struct {
	deps gracefulDependencies
}

func newGraceful(d gracefulDependencies) graceful {
	return graceful{
		deps: gracefulDependencies{
			srv:      d.srv,
			postgres: d.postgres,
			errCh:    d.errCh,
		},
	}
}

func (g graceful) start() {
	if g.deps.postgres != nil {
		err := g.deps.postgres.Close()
		logger.Error("could not close postgres connection", "error", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := g.deps.srv.Shutdown(ctx)
	if err != nil {
		logger.Error("could not shut down server", "error", err)
		g.deps.errCh <- err
	}
	g.deps.errCh <- nil
}
