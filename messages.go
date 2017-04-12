package main

func isStatusByte(c byte) bool {
	return c&0x80 != 0
}
