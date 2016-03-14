package main

import "runtime/debug"

func main() {
	debug.SetTraceback("all")
	
	conf.ReadStart()
	PreloadSlots()
	TerminalSetup()
	l := conf.LastRecalled
	conf.ReadDone()

	RenderSlotToFramebuffer(l)

	StartHTTP()
}
