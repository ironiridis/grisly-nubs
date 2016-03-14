package main

import "time"
import "log"
import "bytes"

func main() {
	var err error
	log.Println("Starting.")

	conf := NewGNConfig()
	err = conf.Read()
	if err != nil {
		log.Println("Configuration file not found. Creating a blank file.")
		conf.Write()
	}

	log.Println("Decoding stored images...")
	slots := make(map[int]*bytes.Buffer)
	for i := 0; i < 10; i++ {
		if conf.Slots[i].Filename == "" {
			log.Printf("Slot %d is empty.\n", i)
		} else {
			log.Printf("Slot %d references file %s.\n", i, conf.Slots[i].Filename)
			slots[i], err = DecodeImageBGR8(conf.Slots[i].Filename)
			if err != nil {
				log.Printf("Slot %d failed decode: %v\n", i, err)
			} else {
				log.Printf("Slot %d decoded.\n", i)
			}
		}
	}

	log.Println("Configuring terminal and starting framebuffer output.")
	TerminalSetup()
	for i := 0; i < 10; i++ {
		b, ok := slots[i]
		if ok {
			err = RenderToFramebuffer(b)
			if err != nil {
				panic(err)
			}
			break
		}
	}

	for {
		<-time.After(1 * time.Hour)
	}
}
