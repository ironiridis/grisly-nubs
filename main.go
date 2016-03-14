package main

func main() {
	conf.ReadStart()
	PreloadSlots()
	TerminalSetup()
	l := conf.LastRecalled
	conf.ReadDone()

	go RenderSlotToFramebuffer(l)
	go ServeTCP()
	StartHTTP()
}
