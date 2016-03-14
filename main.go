package main

func main() {
	conf.ReadStart()
	PreloadSlots()
	TerminalSetup()
	l := conf.LastRecalled
	conf.ReadDone()

	RenderSlotToFramebuffer(l)

	StartHTTP()
}
