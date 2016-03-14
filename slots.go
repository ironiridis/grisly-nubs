package main

import "fmt"
import "log"
import "os"
import "bytes"
import "sync"

type slotMap map[int]*bytes.Buffer

var slots slotMap
var slotsLock sync.RWMutex

// PreloadSlots will do just that.
func PreloadSlots() {
	var err error
	log.Println("Decoding stored images...")
	slots = slotMap{}
	conf.ReadStart()
	for i := 0; i < 10; i++ {
		if conf.Slots[i].Filename == "" {
			log.Printf("Slot %d is empty.\n", i)
		} else {
			err = LoadSlot(i, conf.Slots[i].Filename)
			if err != nil {
				log.Printf("Slot %d failed decode: %v\n", i, err)
			} else {
				log.Printf("Slot %d decoded.\n", i)
			}
		}
	}
	conf.ReadDone()
}

// RenderBufferToFramebuffer takes a pre-decoded image in framebuffer-friendly format
// and writes it out to /dev/fb0.
func RenderBufferToFramebuffer(buf *bytes.Buffer) error {
	fp, err := os.OpenFile("/dev/fb0", os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = fp.Write(buf.Bytes())
	return err
}

// RenderSlotToFramebuffer tries to call RenderBufferToFramebuffer using the cached
// slot data for the slot referenced.
func RenderSlotToFramebuffer(s int) error {
	slotsLock.RLock()
	defer slotsLock.RUnlock()

	b, ok := slots[s]
	if !ok {
		return fmt.Errorf("Cannot load slot %d, appears to be empty", s)
	}

	conf.WriteStart()
	conf.LastRecalled = s
	conf.WriteDone()
	return RenderBufferToFramebuffer(b)
}

func LoadSlot(i int, path string) error {
	var err error
	slotsLock.Lock()
	defer slotsLock.Unlock()
	slots[i], err = DecodeImageBGR8(path)
	return err
}
