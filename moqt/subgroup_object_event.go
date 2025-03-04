package moqt

import (
	"log/slog"

	"github.com/mengelbart/qlog"
)

type SubgroupObjectEventName string

const (
	SubgroupObjectEventCreated SubgroupObjectEventName = "subgroup_object_created"
	SubgroupObjectEventParsed  SubgroupObjectEventName = "subgroup_object_parsed"
)

type SubgroupObjectEvent struct {
	EventName              SubgroupObjectEventName
	StreamID               uint64
	GroupID                uint64
	SubgroupID             uint64
	ObjectID               uint64
	ExtensionHeadersLength uint64
	ExtensionHeaders       ExtensionHeaders
	ObjectPayloadLength    uint64
	ObjectStatus           uint64
	ObjectPayload          qlog.RawInfo
}

func (e SubgroupObjectEvent) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.Uint64("stream_id", e.StreamID),
		slog.Uint64("group_id", e.StreamID),
		slog.Uint64("subgroup_id", e.StreamID),
		slog.Uint64("object_id", e.StreamID),
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

func (e SubgroupObjectEvent) Name() string {
	return string(e.EventName)
}

func (e SubgroupObjectEvent) Category() string {
	return category
}
