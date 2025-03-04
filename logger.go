package qlog

import (
	"context"
	"io"
	"log/slog"
	"time"
)

type jsonSeqWriter struct {
	writer io.Writer
}

func (w *jsonSeqWriter) Write(p []byte) (int, error) {
	return w.writer.Write(append([]byte{'\u001e'}, p...))
}

type Event interface {
	Category() string
	Name() string
	slog.LogValuer
}

type Logger struct {
	logger    *slog.Logger
	reference time.Time
}

func NewQLOGHandler(w io.Writer, title, description, vantagePoint string, schemas ...string) *Logger {
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
	if schemas == nil {
		schemas = []string{}
	}
	logger := slog.New(handler)
	logger.LogAttrs(context.Background(), 0, "",
		slog.String("file_schema", "urn:ietf:params:qlog:file:sequential"),
		slog.String("serialization_format", "application/qlog+json-seq"),
		slog.String("title", title),
		slog.String("description", description),
		slog.Any("event_schemas", schemas),
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
	l.logger.Log(
		context.Background(), slog.LevelInfo,
		e.Category()+":"+e.Name(),
		"data", e,
	)
}
