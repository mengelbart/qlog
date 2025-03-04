package moqt

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/mengelbart/qlog"
	"github.com/mengelbart/qlog/internal/slices"
)

type ExtensionHeaders []ExtensionHeader

func (h ExtensionHeaders) LogValue() slog.Value {
	ehs := slices.Collect(slices.Map(h, func(e ExtensionHeader) json.RawMessage {
		buf, err := json.Marshal(e)
		if err != nil {
			return []byte(fmt.Sprintf("BUG: failed to marshal extension header: %v", err))
		}
		return json.RawMessage(buf)
	}))
	return slog.AnyValue(ehs)
}

type ExtensionHeader struct {
	HeaderType   uint64
	HeaderValue  uint64
	HeaderLength uint64
	Payload      qlog.RawInfo
}

func (h ExtensionHeader) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("header_type", h.HeaderType),
		slog.Uint64("header_value", h.HeaderType),
		slog.Uint64("header_length", h.HeaderType),
		slog.Any("payload", h.Payload),
	)
}

func (h ExtensionHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HeaderType   uint64       `json:"header_type"`
		HeaderValue  uint64       `json:"header_value"`
		HeaderLength uint64       `json:"header_length"`
		Payload      qlog.RawInfo `json:"payload"`
	}{
		HeaderType:   h.HeaderType,
		HeaderValue:  h.HeaderValue,
		HeaderLength: h.HeaderLength,
		Payload:      h.Payload,
	})
}
