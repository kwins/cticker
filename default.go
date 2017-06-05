package cticker

/*
 * 此文件暂时未使用，打算用作持久化存储
 */
import (
	"bytes"
	"encoding/gob"
)

// Handler 定时任务函数
type Handler interface {
	Exec() error
	GobEncoder() ([]byte, error)
	GobDecoder([]byte) error
}
type defaultHandler struct {
}

// Exec Exec
func (dh *defaultHandler) Exec() error {
	return nil
}

// GobEncoder GobEncoder
func (dh *defaultHandler) GobEncoder() ([]byte, error) {
	var buf = bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(dh); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GobDecoder GobDecoder
func (dh *defaultHandler) GobDecoder(b []byte) error {
	var buf = bytes.NewBuffer(nil)
	if _, err := buf.Write(b); err != nil {
		return err
	}
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(dh); err != nil {
		return err
	}
	return nil
}
