package logger

import (
	"context"
	"log/slog"
)

type MultiHandler struct {
	handlers []slog.Handler
}

func (m *MultiHandler) Enabled(
	ctx context.Context,
	l slog.Level,
) bool {

	for _, h := range m.handlers {
		if h.Enabled(ctx, l) {
			return true
		}
	}
	return false
}

func (m *MultiHandler) Handle(
	ctx context.Context,
	r slog.Record,
) error {

	for _, h := range m.handlers {
		rCopy := r.Clone()
		_ = h.Handle(ctx, rCopy)
	}
	return nil
}

func (m *MultiHandler) WithAttrs(
	attrs []slog.Attr,
) slog.Handler {

	hs := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		hs[i] = h.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: hs}
}

func (m *MultiHandler) WithGroup(
	name string,
) slog.Handler {

	hs := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		hs[i] = h.WithGroup(name)
	}
	return &MultiHandler{handlers: hs}
}
