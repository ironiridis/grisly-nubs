package main

import "image/png"
import "image/jpeg"
import "bytes"
import "os"

func decodePNG(path string) (bytes.Buffer, error) {
	fp, err := os.Open(path)
	defer fp.Close()
	if err != nil {
		return nil, err
	}
	im, err := png.Decode(fp)
	if err != nil {
		return nil, err
	}
	return imageToBGRBuffer(im)
}

func decodeJPEG(path string) (bytes.Buffer, error) {
	fp, err := os.Open(path)
	defer fp.Close()
	if err != nil {
		return nil, err
	}
	im, err := jpeg.Decode(fp)
	if err != nil {
		return nil, err
	}
	return imageToBGRBuffer(im)
}

func imageToBGRBuffer(im *image.Image) (buf bytes.Buffer, err error) {
	for y := im.Bounds().Min.Y; y < im.Bounds().Max.Y; y++ {
		for x := im.Bounds().Min.X; x < im.Bounds().Max.X; x++ {
			pr, pg, pb, _ := im.At(x, y).RGBA()
			buf.Write([]byte{byte(pb >> 8), byte(pg >> 8), byte(pr >> 8)})
		}
	}
	return
}
