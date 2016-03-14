package main

import "io"
import "bytes"
import "os"
import "runtime"

func CrashHard() {
	b := new(bytes.Buffer)
	b.Grow(3145728) // 3mb
	runtime.Stack(b.Bytes(), true)
	
	fp, _ := os.Create("/home/charrington/diag.txt")
	io.Copy(fp, b)
	fp.Close()
	
}
