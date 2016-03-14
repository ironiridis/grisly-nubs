package main

import "os/exec"

// CmdRemountRO remounts the root filesystem read-only.
var CmdRemountRO *exec.Cmd

// CmdRemountRW remounts the root filesystem read-write. This should ideally be brief.
var CmdRemountRW *exec.Cmd

func init() {
	CmdRemountRO = exec.Command("/bin/mount", "-n", "-o", "remount,ro", "/")
	CmdRemountRW = exec.Command("/bin/mount", "-n", "-o", "remount,rw", "/")
}
