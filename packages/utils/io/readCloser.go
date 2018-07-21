package io

import (
	"bytes"
	"io"
)

// ReadCloserToString turns the stream into a string
func ReadCloserToString(rc io.ReadCloser) (res string, err error) {
	buf, err := readCloserToBuffer(rc)
	if err != nil {
		return
	}
	res = buf.String()
	return
}

//ReadCloserToByte turns the stream into a byte array
func ReadCloserToByte(rc io.ReadCloser) (res []byte, err error) {
	buf, err := readCloserToBuffer(rc)
	if err != nil {
		return
	}
	res = buf.Bytes()
	return
}

func readCloserToBuffer(rc io.ReadCloser) (buf *bytes.Buffer, err error) {
	buf = new(bytes.Buffer)
	_, err = buf.ReadFrom(rc)
	defer rc.Close()
	return
}
