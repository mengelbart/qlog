package moqt

import (
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
	Value  qlog.RawInfo
}
