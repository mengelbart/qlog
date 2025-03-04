package moqt

import "log/slog"

type SubgroupHeaderEventName string

const (
	SubgroupHeaderEventCreated SubgroupHeaderEventName = "subgroup_header_created"
	SubgroupHeaderEventParsed  SubgroupHeaderEventName = "subgroup_header_parsed"
)

type SubgroupHeaderEvent struct {
	EventName         SubgroupHeaderEventName
	StreamID          uint64
	TrackAlias        uint64
	GroupID           uint64
	SubgroupID        uint64
	PublisherPriority uint8
}

func (e SubgroupHeaderEvent) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("stream_id", e.StreamID),
		slog.Uint64("track_alias", e.TrackAlias),
		slog.Uint64("group_id", e.GroupID),
		slog.Uint64("subgroup_id", e.SubgroupID),
		slog.Uint64("publisher_priority", uint64(e.PublisherPriority)),
	)
}

func (e SubgroupHeaderEvent) Name() string {
	return string(e.EventName)
}

func (e SubgroupHeaderEvent) Category() string {
	return category
}
