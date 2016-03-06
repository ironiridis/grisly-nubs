package main

import "os"
import "bytes"

func renderToFramebuffer(buf *bytes.Buffer) error {
	fp, err := os.OpenFile("/dev/fb0", os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = buf.WriteTo(fp)
	return err
}
