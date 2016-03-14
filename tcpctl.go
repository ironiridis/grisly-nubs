package main

import "net"
import "text/scanner"

func TCPSession(s net.Conn) {
	var l scanner.Scanner
	l.Init(s)
	var tok rune
	for tok != scanner.EOF {
		tok = l.Scan()
		switch l.TokenText() {
		case "slot0":
			RenderSlotToFramebuffer(0)
		case "slot1":
			RenderSlotToFramebuffer(1)
		case "slot2":
			RenderSlotToFramebuffer(2)
		case "slot3":
			RenderSlotToFramebuffer(3)
		case "slot4":
			RenderSlotToFramebuffer(4)
		case "slot5":
			RenderSlotToFramebuffer(5)
		case "slot6":
			RenderSlotToFramebuffer(6)
		case "slot7":
			RenderSlotToFramebuffer(7)
		case "slot8":
			RenderSlotToFramebuffer(8)
		case "slot9":
			RenderSlotToFramebuffer(9)

		}
	}
}

func ServeTCP() {
	sockListen, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	for {
		sock, err := sockListen.Accept()
		if err != nil {
			// well, ...
		} else {
			go TCPSession(sock)
		}
	}
}
