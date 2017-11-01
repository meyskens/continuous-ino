package storage

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Run is the format to store a run in
type Run struct {
	ID      uint64      `json:"id"`
	Repo    string      `json:"repo"`
	SHA     string      `json:"sha"`
	Time    time.Time   `json:"time"`
	Output  []RunOutput `json:"output"`
	Errors  []string    `json:"errors"`
	Running bool        `json:"running"`
}

// RunOutput is the output of a specific run of a test
type RunOutput struct {
	Name   string `json:"name"`
	File   string `json:"file"`
	Step   string `json:"step"`
	Output string `json:"output"`
}

// Encode encodes a run to gob
func (r *Run) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeRun decodes gob to Run
func DecodeRun(data []byte) (Run, error) {
	var r Run
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&r)
	if err != nil {
		return Run{}, err
	}
	return r, nil
}
