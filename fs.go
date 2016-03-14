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

func FSSetBootProd() {
	exec.Command("/bin/mkdir", "-p", "/boot").Run()
	exec.Command("/bin/mount", "-n", "/boot").Run()
	exec.Command("/bin/cp", "/boot/_bootprod.cmdline.txt", "/boot/cmdline.txt").Run()
	exec.Command("/bin/umount", "-n", "/boot").Run()
}

func FSSetBootDev() {
	exec.Command("/bin/mkdir", "-p", "/boot").Run()
	exec.Command("/bin/mount", "-n", "/boot").Run()
	exec.Command("/bin/cp", "/boot/_bootdev.cmdline.txt", "/boot/cmdline.txt").Run()
	exec.Command("/bin/umount", "-n", "/boot").Run()
}

func init() {
	// we need the slots directory to store our assets
	exec.Command("/bin/mkdir", "-p", "/slots").Run()
}
