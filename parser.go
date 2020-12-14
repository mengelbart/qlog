package qlog

import (
	"io"

	"github.com/francoispqt/gojay"
)

func Parse(r io.Reader) (*QLOGFile, error) {
	dec := gojay.BorrowDecoder(r)
	defer dec.Release()
	qf := &QLOGFile{}
	err := dec.Decode(qf)
	return qf, err
}
