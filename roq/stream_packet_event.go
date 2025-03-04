package roq

import "log/slog"

type StreamPacketEventName = string

const (
	StreamPacketEventTypeCreated StreamPacketEventName = "stream_packet_created"
	StreamPacketEventTypeParsed  StreamPacketEventName = "stream_packet_parsed"
)

type StreamPacketEvent struct {
	EventName StreamPacketEventName
	StreamID  int64
	Packet    Packet
}

func (e StreamPacketEvent) Category() string {
	return roqCategory
}

func (e StreamPacketEvent) Name() string {
	return e.EventName
}

func (e StreamPacketEvent) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int64("stream_id", e.StreamID),
		slog.Attr{
			Key:   "packet",
			Value: e.Packet.LogValue(),
		},
	)
}
