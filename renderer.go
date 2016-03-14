package main

import "os"
import "bytes"

// RenderToFramebuffer takes a pre-decoded image in framebuffer-friendly format
// and writes it out to /dev/fb0.
func RenderToFramebuffer(buf *bytes.Buffer) error {
	fp, err := os.OpenFile("/dev/fb0", os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = buf.WriteTo(fp)
	return err
}
