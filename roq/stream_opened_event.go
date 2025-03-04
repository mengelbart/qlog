package roq

import "log/slog"

type StreamOpenedEvent struct {
	FlowID   uint64
	StreamID uint64
}

func (e StreamOpenedEvent) Category() string {
	return roqCategory
}

func (e StreamOpenedEvent) Name() string {
	return "stream_opened"
}

func (e StreamOpenedEvent) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("flow_id", e.FlowID),
		slog.Uint64("stream_id", e.StreamID),
	)
}
