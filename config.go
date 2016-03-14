package main

import "log"
import "os"
import "sync"
import "encoding/json"

// GNConfigSlot contains the label and filename for each image slot.
type GNConfigSlot struct {
	Label, Filename string
}

// GNConfig is a container for GNConfigSlot elements, and the last recalled slot
// number.
type GNConfig struct {
	Slots        [10]GNConfigSlot
	LastRecalled int
}

var conf GNConfig
var conflock sync.RWMutex

func init() {
	err := conf.Read()
	if err != nil {
		log.Println("Configuration file not found. Creating a blank file.")
		conf.Write()
	}
}

// Read attempts to pull our configuration data from storage. It will return an error if our
// configuration file is missing, and panic if it looks like it is corrupt.
func (c *GNConfig) Read() error {
	fp, err := os.Open("/gn_config.json")
	if err != nil {
		return err
	}
	defer fp.Close()

	jdec := json.NewDecoder(fp)
	err = jdec.Decode(c)
	if err != nil {
		panic(err)
	}

	return nil
}

// Write will temporarily make our disk writable, flush out our configuration data, and revert
// back to read-only. It will panic if anything goes wrong, which can happen.
func (c *GNConfig) Write() {
	err := CmdRemountRW.Run()
	if err != nil {
		panic(err)
	}

	fp, err := os.Create("/gn_config.json")
	if err != nil {
		panic(err)
	}

	jenc := json.NewEncoder(fp)
	err = jenc.Encode(&c)
	if err != nil {
		panic(err)
	}

	err = fp.Close()
	if err != nil {
		panic(err)
	}

	err = CmdRemountRO.Run()
	if err != nil {
		panic(err)
	}

	return
}

func (c *GNConfig) ReadStart() {
	conflock.RLock()
}

func (c *GNConfig) ReadDone() {
	conflock.RUnlock()
}

func (c *GNConfig) WriteStart() {
	conflock.Lock()
}

func (c *GNConfig) WriteDone() {
	c.Write()
	conflock.Unlock()
}
