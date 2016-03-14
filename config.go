package main

import "os"
import "os/exec"
import "encoding/json"

// GNConfigSlot contains the label and filename for each image slot.
type GNConfigSlot struct {
	Label, Filename string
}

// GNConfig is a container for GNConfigSlot elements. It may also contain some other
// values in the future.
type GNConfig struct {
	Slots [10]GNConfigSlot
}

// NewGNConfig returns a mint condition GNConfig.
func NewGNConfig() *GNConfig {
	return &GNConfig{}
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

	fp, err := os.Open("/gn_config.json")
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
