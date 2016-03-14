package main

import "os"

// TerminalSetup writes a series of escape sequences to os.Stdout to set up the display
// in the most ideal way.
func TerminalSetup() {
	os.Stdout.WriteString("\x1b[9;0]")         // blank 0
	os.Stdout.WriteString("\x1b[14;0]")        // powerdown 0
	os.Stdout.WriteString("\x1b[41m")          // red background (for troubleshooting)
	os.Stdout.WriteString("\x1b[?25l\x1b[?1c") // disable cursor
	os.Stdout.WriteString("\x1b[8]")           // store defaults
	os.Stdout.WriteString("\x1b[H\x1b[J")      // clear screen
}
