package roq

import (
	"log/slog"

	"github.com/mengelbart/qlog"
)

type Packet struct {
	FlowID uint64
	Length uint64
	Raw    *qlog.RawInfo
}

func (p *Packet) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.Uint64("flow_id", p.FlowID),
		slog.Uint64("length", p.Length),
	}
	if p.Raw != nil {
		slog.Any("raw", p.Raw)
	}
	return slog.GroupValue(attrs...)
}
