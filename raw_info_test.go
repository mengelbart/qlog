package qlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawInfo(t *testing.T) {
	ri := RawInfo{
		Length:        1,
		PayloadLength: 2,
		Data:          []byte("abc"),
	}
	buf, err := ri.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"length":1,"payload_length":2,"data":"616263"}`, string(buf))
}
