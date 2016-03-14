package main

import "fmt"
import "io"
import "image"
import _ "image/png"
import _ "image/jpeg"
import "bytes"
import "os"

// DecodeImageBGR8 reads a supported image, specified by path, and stores it in a freshly-
// allocated bytes.Buffer, which it returns.
func DecodeImageBGR8(path string) (*bytes.Buffer, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	im, _, err := image.Decode(fp)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	// pre-allocate our storage
	buf.Grow(3 * im.Bounds().Dx() * im.Bounds().Dy())
	for y := im.Bounds().Min.Y; y < im.Bounds().Max.Y; y++ {
		for x := im.Bounds().Min.X; x < im.Bounds().Max.X; x++ {
			pr, pg, pb, _ := im.At(x, y).RGBA()
			buf.Write([]byte{byte(pb >> 8), byte(pg >> 8), byte(pr >> 8)})
		}
	}
	return buf, nil
}

// ValidImage tries to decode the image and does some sanity checking of the image file.
// It returns a boolean for validity, and an error if relevant.
func ValidImage(r io.Reader) (bool, error) {
	c, _, err := image.DecodeConfig(r)
	if err != nil {
		return false, err
	}
	if c.Width != 1920 || c.Height != 1080 {
		return false, fmt.Errorf("Image must be exactly 1920x1080; this image is %dx%d.", c.Width, c.Height)
	}
	return true, nil
}
