package main

func main() {
	var err error
	conf.ReadStart()
	PreloadSlots()
	conf.ReadDone()
	TerminalSetup()
	for i := 0; i < 10; i++ {
		_, ok := slots[i]
		if ok {
			err = RenderSlotToFramebuffer(i)
			if err != nil {
				panic(err)
			}
			break
		}
	}

	StartHTTP()
}
