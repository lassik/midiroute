package main

import (
	"fmt"
	"os"
)

// #cgo LDFLAGS: -framework Foundation -framework CoreMIDI
// extern void darwinCoreLoop(void);
// int isatty(int fd);
import "C"

//export darwinReceiveMidiByte
func darwinReceiveMidiByte(c byte) {
	receiveMidiByte(c)
}

func isTerminalFd(fd uintptr) bool {
	return C.isatty(C.int(fd)) != 0
}

func main() {
	if err := setup(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	C.darwinCoreLoop()
}
