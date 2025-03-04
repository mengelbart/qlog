package moqt

import (
	"log/slog"

	"github.com/mengelbart/qlog"
)

type ParameterName string

const (
	ParameterNameAuthorizationInfo = "authorization_info"
	ParameterNameDeliveryTimeout   = "delivery_timeout"
	ParameterNameMaxCacheDuration  = "max_cache_duration"
	ParameterNamePath              = "path"
	ParameterNameMaxSubscribeID    = "max_subscribe_id"
)

type Parameter struct {
	Name   ParameterName
	Length uint64
	Value  *qlog.RawInfo
}

func (p Parameter) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.String("name", string(p.Name)),
		slog.Uint64("length", p.Length),
	}
	if p.Value != nil {
		attrs = append(attrs, slog.Any("value", p.Value))
	}
	return slog.GroupValue(attrs...)
}
