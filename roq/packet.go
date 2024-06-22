package roq

import "log/slog"

type Packet struct {
	FlowID uint64
	Length int
}

func (p *Packet) attrs() []slog.Attr {
	return []slog.Attr{
		slog.Uint64("flow_id", p.FlowID),
		slog.Int("length", p.Length),
	}
}
