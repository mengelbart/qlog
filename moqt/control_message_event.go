package moqt

import (
	"log/slog"

	"github.com/mengelbart/qlog"
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
	Raw       *qlog.RawInfo
}

func (e ControlMessageEvent) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.Uint64("stream_id", e.StreamID),
	}
	if e.Length > 0 {
		attrs = append(attrs, slog.Uint64("length", e.Length))
	}
	attrs = append(attrs, slog.Attr{
		Key:   "message",
		Value: e.Message.LogValue(),
	})
	if e.Raw != nil {
		attrs = append(attrs, slog.Any("raw", e.Raw))
	}
	return slog.GroupValue(attrs...)
}

// Name implements qlog.Event.
func (e ControlMessageEvent) Name() string {
	return string(e.EventName)
}

// Name implements qlog.Event.
func (e ControlMessageEvent) Category() string {
	return category
}
