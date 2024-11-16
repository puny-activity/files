package lggr

import (
	"context"
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	actionID := ctx.Value("actionId")
	if actionID != nil {
		r.AddAttrs(slog.String("actionId", actionID.(string)))
	}
	return h.Handler.Handle(ctx, r)
}
