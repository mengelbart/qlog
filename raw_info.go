package qlog

import (
	"encoding/hex"
	"encoding/json"
	"log/slog"
)

type RawInfo struct {
	Length        uint64
	PayloadLength uint64
	Data          []byte
}

func (i RawInfo) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("length", i.Length),
		slog.Uint64("payload_length", i.PayloadLength),
		slog.Any("data", hex.EncodeToString(i.Data)),
	)
}

func (i RawInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Length        uint64 `json:"length"`
		PayloadLength uint64 `json:"payload_length"`
		Data          string `json:"data"`
	}{
		Length:        i.Length,
		PayloadLength: i.PayloadLength,
		Data:          hex.EncodeToString(i.Data),
	})
}
