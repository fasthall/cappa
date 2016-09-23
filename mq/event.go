package mq

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Event struct {
	Time   time.Time
	Type   string
	Action string
	Bucket string
	Object string
}

func (e *Event) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(e)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(e []byte) (Event, error) {
	buf := bytes.NewBuffer(e)
	dec := gob.NewDecoder(buf)
	var to Event
	err := dec.Decode(&to)
	return to, err
}
