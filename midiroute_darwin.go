package main

// #cgo LDFLAGS: -framework Foundation -framework CoreMIDI
// void darwinCoreLoop();
import "C"

//export darwinReceiveMidiByte
func darwinReceiveMidiByte(byt int) {
	receiveMidiByte(byt)
}

func main() {
	C.darwinCoreLoop()
}
