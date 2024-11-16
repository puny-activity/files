package lggr

import (
	"context"
	"github.com/puny-activity/files/pkg/base"
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	actionID := ctx.Value(base.ActionID)
	if actionID != nil {
		r.AddAttrs(slog.String(base.ActionID, actionID.(string)))
	}
	return h.Handler.Handle(ctx, r)
}
