package moqt

import (
	"log/slog"
)

type ControlMessageEventName string

const (
	ControlMessageEventCreated ControlMessageEventName = "control_message_created"
	ControlMessageEventParsed  ControlMessageEventName = "control_message_parsed"
)

type ControlMessageEvent struct {
	EventName ControlMessageEventName
	StreamID  uint64
	Length    uint64
	Message   slog.LogValuer
}

func (e ControlMessageEvent) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("stream_id", e.StreamID),
		slog.Uint64("length", e.Length),
		slog.Attr{
			Key:   "message",
			Value: e.Message.LogValue(),
		},
		// slog.Group("message", slices.Collect(slices.Map(e.Message.Attrs(), func(e slog.Attr) any { return e }))...),
	)
}

// Name implements qlog.Event.
func (e ControlMessageEvent) Name() string {
	return string(e.EventName)
}

// Name implements qlog.Event.
func (e ControlMessageEvent) Category() string {
	return category
}
