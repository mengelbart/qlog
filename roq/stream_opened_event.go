package roq

import "log/slog"

type StreamOpenedEvent struct {
	FlowID   uint64
	StreamID int64
}

func (e StreamOpenedEvent) Category() string {
	return roqCategory
}

func (e StreamOpenedEvent) Name() string {
	return "stream_opened"
}

func (e StreamOpenedEvent) Attrs() []slog.Attr {
	return []slog.Attr{
		slog.Uint64("flow_id", e.FlowID),
		slog.Int64("stream_id", e.StreamID),
	}
}
