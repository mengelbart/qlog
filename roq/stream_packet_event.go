package roq

import "log/slog"

type StreamPacketEventType = string

const (
	StreamPacketEventTypeCreated StreamPacketEventType = "stream_packet_created"
	StreamPacketEventTypeParsed  StreamPacketEventType = "stream_packet_parsed"
)

type StreamPacketEvent struct {
	Type     StreamPacketEventType
	StreamID int64
	Packet   Packet
}

func (e StreamPacketEvent) Category() string {
	return roqCategory
}

func (e StreamPacketEvent) Name() string {
	return e.Type
}

func (e StreamPacketEvent) Attrs() []slog.Attr {
	return append(
		e.Packet.attrs(),
		slog.Int64("stream_id", e.StreamID),
	)
}
