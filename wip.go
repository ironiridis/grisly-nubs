package main

import "fmt"
import "time"

func main() {
	// set up terminal interface
	fmt.Print("\x1b[9;0]")         // blank 0
	fmt.Print("\x1b[14;0]")        // powerdown 0
	fmt.Print("\x1b[41m")          // red background (for troubleshooting)
	fmt.Print("\x1b[?25l\x1b[?1c") // disable cursor
	fmt.Print("\x1b[8]")           // store defaults
	fmt.Print("\x1b[H\x1b[J")      // clear screen

	b, err := decodeImageBGR8("/home/charrington/1.jpg")
	if err != nil {
		panic(err)
	}
	err = renderToFramebuffer(b)
	if err != nil {
		panic(err)
	}

	for {
		<-time.After(1 * time.Hour)
	}
}
