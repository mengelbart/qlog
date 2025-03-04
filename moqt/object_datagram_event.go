package moqt

import (
	"log/slog"

	"github.com/mengelbart/qlog"
)

type ObjectDatagramEventName string

const (
	ObjectDatagramEventCreated       ObjectDatagramEventName = "object_datagram_created"
	ObjectDatagramEventparsed        ObjectDatagramEventName = "object_datagram_parsed"
	ObjectDatagramStatusEventCreated ObjectDatagramEventName = "object_datagram_status_created"
	ObjectDatagramStatusEventparsed  ObjectDatagramEventName = "object_datagram_status_parsed"
)

type ObjectDatagramEvent struct {
	EventName              ObjectDatagramEventName
	TrackAlias             uint64
	GroupID                uint64
	ObjectID               uint64
	PublisherPriority      uint8
	ExtensionHeadersLength uint64
	ExtensionHeaders       ExtensionHeaders
	ObjectStatus           uint64
	Payload                qlog.RawInfo
}

func (e ObjectDatagramEvent) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.Uint64("track_alias", e.TrackAlias),
		slog.Uint64("group_id", e.GroupID),
		slog.Uint64("object_id", e.ObjectID),
		slog.Any("publisher_priority", e.PublisherPriority),
		slog.Uint64("extension_headers_length", e.ExtensionHeadersLength),
	}
	if len(e.ExtensionHeaders) > 0 {
		attrs = append(attrs, slog.Any("extension_headers", e.ExtensionHeaders))
	}
	if e.EventName == ObjectDatagramStatusEventCreated || e.EventName == ObjectDatagramStatusEventparsed {
		attrs = append(attrs, slog.Uint64("object_status", e.ObjectStatus))
	}
	if e.EventName == ObjectDatagramEventCreated || e.EventName == ObjectDatagramEventparsed {
		attrs = append(attrs, slog.Any("object_payload", e.Payload))
	}
	return slog.GroupValue(attrs...)
}

func (e ObjectDatagramEvent) Name() string {
	return string(e.EventName)
}

func (e ObjectDatagramEvent) Category() string {
	return category
}
