package utils

import (
	"io"
	"bytes"
)

func IoRCTS(reader io.ReadCloser) (str string, err error) {
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(reader)
	if err != nil {
		return
	}

	str = buffer.String()
	return
}
