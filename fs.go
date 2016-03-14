package main

import "os/exec"

// FSRemountRO remounts the root filesystem read-only.
func FSRemountRO() {
	exec.Command("/bin/mount", "-n", "-o", "remount,ro", "/").Run()
}

// FSRemountRW remounts the root filesystem read-write. This should ideally be brief.
func FSRemountRW() {
	exec.Command("/bin/mount", "-n", "-o", "remount,rw", "/").Run()
}

func init() {
	// we need the slots directory to store our assets
	exec.Command("/bin/mkdir", "-p", "/slots").Run()
}
