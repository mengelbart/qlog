package roq

import "log/slog"

type DatagramPacketEventType = string

const (
	DatagramPacketEventTypeCreated DatagramPacketEventType = "datagram_packet_created"
	DatagramPacketEventTypeParsed  DatagramPacketEventType = "datagram_packet_parsed"
)

type DatagramPacketEvent struct {
	Type   DatagramPacketEventType
	Packet Packet
}

func (e DatagramPacketEvent) Category() string {
	return roqCategory
}

func (e DatagramPacketEvent) Name() string {
	return e.Type
}

func (e DatagramPacketEvent) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Attr{
			Key:   "packet",
			Value: e.Packet.LogValue(),
		},
	)
}
