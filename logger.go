package qlog

import (
	"io"
	"log/slog"
	"time"

	"github.com/mengelbart/qlog/roq"
)

type Event interface {
	Category() string
	Name() string
	Attrs() []slog.Attr
}

type Logger struct {
	logger    *slog.Logger
	reference time.Time
}

type jsonSeqWriter struct {
	writer io.Writer
}

func (w *jsonSeqWriter) Write(p []byte) (int, error) {
	return w.writer.Write(append([]byte{'\u001e'}, p...))
}

func NewQLOGHandler(w io.Writer, title, vantagePoint string) *Logger {
	reference := time.Now()
	initTime := false
	initName := false
	handler := slog.NewJSONHandler(&jsonSeqWriter{writer: w}, &slog.HandlerOptions{
		AddSource: false,
		Level:     nil,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "msg":
				if !initName {
					initName = true
					return slog.Attr{}
				}
				return slog.Attr{
					Key:   "name",
					Value: a.Value,
				}
			case "level":
				return slog.Attr{}
			case "time":
				if !initTime {
					initTime = true
					return slog.Attr{}
				}
				d := a.Value.Time().Sub(reference)
				return slog.Float64("time", float64(d.Nanoseconds())/1e6)
			}
			return a
		},
	})
	logger := slog.New(handler)
	logger.LogAttrs(nil, 0, "",
		slog.String("qlog_version", "draft-02"),
		slog.String("qlog_format", "JSON-SEQ"),
		slog.String("title", title),
		slog.Group("trace",
			slog.Group("vantage_point", slog.String("type", vantagePoint)),
			slog.Group("common_fields",
				slog.String("time_format", "relative"),
				slog.Float64("reference_time", float64(reference.UnixNano())/1e6),
			),
		),
	)
	return &Logger{
		logger:    logger,
		reference: reference,
	}
}

func (l *Logger) Log(e Event) {
	anys := []any{}
	for _, a := range e.Attrs() {
		anys = append(anys, a)
	}
	l.logger.LogAttrs(nil, 0,
		e.Category()+":"+e.Name(),
		slog.Group("data", anys...),
		// e.Attrs()...,
	)
}

func (l *Logger) RoQStreamOpened(flowID uint64, streamID int64) {
	l.Log(roq.StreamOpenedEvent{
		FlowID:   flowID,
		StreamID: streamID,
	})
}

func (l *Logger) RoQStreamPacketCreated(flowID uint64, streamID int64, length int) {
	l.Log(roq.StreamPacketEvent{
		Type:     roq.StreamPacketEventTypeCreated,
		StreamID: streamID,
		Packet: roq.Packet{
			FlowID: flowID,
			Length: length,
		},
	})
}

func (l *Logger) RoQStreamPacketParsed(flowID uint64, streamID int64, length int) {
	l.Log(roq.StreamPacketEvent{
		Type:     roq.StreamPacketEventTypeParsed,
		StreamID: streamID,
		Packet: roq.Packet{
			FlowID: flowID,
			Length: length,
		},
	})
}

func (l *Logger) RoQDatagramPacketCreated(flowID uint64, length int) {
	l.Log(roq.DatagramPacketEvent{
		Type: roq.DatagramPacketEventTypeCreated,
		Packet: roq.Packet{
			FlowID: flowID,
			Length: length,
		},
	})
}

func (l *Logger) RoQDatagramPacketParsed(flowID uint64, length int) {
	l.Log(roq.DatagramPacketEvent{
		Type: roq.DatagramPacketEventTypeParsed,
		Packet: roq.Packet{
			FlowID: flowID,
			Length: length,
		},
	})
}
