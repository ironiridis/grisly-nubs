package main

func main() {
	b, err := decodeImageBGR8("/home/charrington/1.jpg")
	if err != nil {
		panic(err)
	}
	err = renderToFramebuffer(b)
	if err != nil {
		panic(err)
	}
}
