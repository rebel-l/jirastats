package io

import (
	"bytes"
	"io"
)

func ReadCloserToString(rc io.ReadCloser) (res string, err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(rc)
	if err != nil {
		return
	}
	defer rc.Close()

	res = buf.String()
	return
}
