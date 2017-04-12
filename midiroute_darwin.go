package main

import (
	"fmt"
	"os"
)

// #cgo LDFLAGS: -framework Foundation -framework CoreMIDI
// extern void darwinCoreLoop(void);
import "C"

//export darwinReceiveMidiByte
func darwinReceiveMidiByte(byt int) {
	receiveMidiByte(byt)
}

func main() {
	if err := setup(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	C.darwinCoreLoop()
}
