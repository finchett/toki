package main

import (
	"context"

	"github.com/go-faster/errors"

	"github.com/kriuchkov/tock/internal/app/runtime"
	"github.com/kriuchkov/tock/internal/core/models"
)

// App is the Wails-bound application surface. It owns a tock Runtime so the
// desktop UI calls the same services the `tock` CLI does.
type App struct {
	ctx context.Context
	rt  *runtime.Runtime
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	rt, err := runtime.Load(ctx, runtime.Request{})
	if err != nil {
		// Surface as nil rt; methods below will return a typed error to the UI.
		return
	}
	a.rt = rt
}

// ListRecent returns up to `limit` most recent activities. The hello-world
// integration point: it proves the desktop app talks to tock's services.
func (a *App) ListRecent(limit int) ([]models.Activity, error) {
	if a.rt == nil {
		return nil, errors.New("tock runtime is not initialized")
	}
	if limit <= 0 {
		limit = 20
	}
	return a.rt.ActivityService.GetRecent(a.ctx, limit)
}
