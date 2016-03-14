package main

import "os"
import "os/exec"
import "encoding/json"

type GNSlot struct {
	Label, Filename string
}

type GNConfig struct {
	Slots []GNSlot
}

var Config GNConfig

var CmdRemountRO *exec.Cmd
var CmdRemountRW *exec.Cmd

func init() {
	CmdRemountRO = exec.Command("/bin/mount", "-n", "-o", "remount,ro", "/")
	CmdRemountRW = exec.Command("/bin/mount", "-n", "-o", "remount,rw", "/")
}

func (c *GNConfig) Read() {
	fp, err := os.Open("/gn_config.json")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	jdec := json.NewDecoder(fp)
	err = jdec.Decode(c)
	if err != nil {
		panic(err)
	}

	return
}

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
