package moqt

import (
	"log/slog"

	"github.com/mengelbart/qlog"
)

type FetchObjectEventName string

const (
	FetchObjectEventCreated FetchObjectEventName = "fetch_object_created"
	FetchObjectEventParsed  FetchObjectEventName = "fetch_object_parsed"
)

type FetchObjectEvent struct {
	EventName              FetchObjectEventName
	StreamID               uint64
	GroupID                uint64
	SubgroupID             uint64
	ObjectID               uint64
	PublisherPriority      uint8
	ExtensionHeadersLength uint64
	ExtensionHeaders       ExtensionHeaders
	ObjectPayloadLength    uint64
	ObjectStatus           uint64
	ObjectPayload          qlog.RawInfo
}

func (e FetchObjectEvent) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.Uint64("stream_id", e.StreamID),
		slog.Uint64("group_id", e.StreamID),
		slog.Uint64("subgroup_id", e.StreamID),
		slog.Uint64("object_id", e.StreamID),
		slog.Any("publisher_priority", e.PublisherPriority),
		slog.Uint64("extension_headers_length", e.ExtensionHeadersLength),
		slog.Any("extension_headers", e.ExtensionHeaders),
	}
	if e.ObjectPayloadLength == 0 {
		attrs = append(attrs, slog.Uint64("object_status", e.ObjectStatus))
	}
	if e.ObjectPayloadLength > 0 {
		attrs = append(attrs, slog.Any("object_payload", e.ObjectPayload))
	}
	return slog.GroupValue(attrs...)
}

func (e FetchObjectEvent) Name() string {
	return string(e.EventName)
}

func (e FetchObjectEvent) Category() string {
	return category
}
