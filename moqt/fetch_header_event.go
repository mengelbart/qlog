package moqt

import "log/slog"

type FetchHeaderEventName string

const (
	FetchHeaderEventCreated FetchHeaderEventName = "fetch_header_created"
	FetchHeaderEventParsed  FetchHeaderEventName = "fetch_header_parsed"
)

type FetchHeaderEvent struct {
	EventName   FetchHeaderEventName
	StreamID    uint64
	SubscribeID uint64
}

func (e FetchHeaderEvent) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("stream_id", e.StreamID),
		slog.Uint64("subscribe_id", e.SubscribeID),
	)
}

func (e FetchHeaderEvent) Name() string {
	return string(e.EventName)
}

func (e FetchHeaderEvent) Category() string {
	return category
}
