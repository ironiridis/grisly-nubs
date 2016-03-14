package main

import "time"

func main() {
	TerminalSetup()
	b, err := DecodeImageBGR8("/home/charrington/1.jpg")
	if err != nil {
		panic(err)
	}
	err = RenderToFramebuffer(b)
	if err != nil {
		panic(err)
	}

	for {
		<-time.After(1 * time.Hour)
	}
}
