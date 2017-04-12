package main

const (
	NoteOff = 0x80
	NoteOn  = 0x90
)

func clampToByte(c int) byte {
	if c < 0 {
		return 0
	}
	if c > 127 {
		return 127
	}
	return byte(c)
}

func isStatusByte(c byte) bool {
	return c&0x80 != 0
}

func isMsg(m []byte, wantedMsg byte) bool {
	return len(m) >= 2 && (m[0]&0xf0 == wantedMsg)
}
