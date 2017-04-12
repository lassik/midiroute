package main

import "fmt"

func receiveMidiByte(byt int) {
	fmt.Printf("CORE got midi byte 0x%02x\n", byt)
}
