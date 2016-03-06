package main

import "image/png"
import "flag"
import "os"

func main() {
	loadpng := flag.String("f", "test.png", "file to load to the framebuffer")
	flag.Parse()

	f, err := os.Open(*loadpng)
	if err != nil {
		panic(err)
	}
	i, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	f.Close()
	fb, err := os.OpenFile("/dev/fb0", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	for y := i.Bounds().Min.Y; y < i.Bounds().Max.Y; y++ {
		for x := i.Bounds().Min.X; x < i.Bounds().Max.X; x++ {
			pr, pg, pb, _ := i.At(x, y).RGBA()
			fb.Write([]byte{byte(pb >> 8), byte(pg >> 8), byte(pr >> 8), byte(0)})
		}
	}
}


