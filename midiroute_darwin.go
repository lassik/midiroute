package main

// #cgo LDFLAGS: -framework Foundation -framework CoreMIDI
// extern void darwinCoreLoop(void);
import "C"

//export darwinReceiveMidiByte
func darwinReceiveMidiByte(byt int) {
	receiveMidiByte(byt)
}

func main() {
	C.darwinCoreLoop()
}
