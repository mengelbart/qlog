package moqt

import "log/slog"

type StreamType string

const (
	StreamTypeSubgroupHeader StreamType = "subgroup_header"
	StreamTypeFetchHeader    StreamType = "fetch_header"
)

type StreamTypeSetEvent struct {
	Owner      *Owner
	StreamID   uint64
	StreamType StreamType
}

func (e StreamTypeSetEvent) LogValue() slog.Value {
	attrs := []slog.Attr{}
	if e.Owner != nil {
		attrs = append(attrs, slog.String("owner", string(*e.Owner)))
	}
	attrs = append(attrs,
		slog.Uint64("stream_id", e.StreamID),
		slog.String("stream_type", string(e.StreamType)),
	)
	return slog.GroupValue(attrs...)
}

func (e StreamTypeSetEvent) Name() string {
	return "stream_type_set"
}

func (e StreamTypeSetEvent) Category() string {
	return category
}
