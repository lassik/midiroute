package main

import (
	"bytes"
	"fmt"
)

func msgToMsg(msg []byte) []byte {
	return msg
}

func msgToDump(msg []byte) []byte {
	var buf bytes.Buffer
	delim := ' '
	for i, byt := range msg {
		if i == len(msg)-1 {
			delim = '\n'
		}
		buf.WriteString(fmt.Sprintf("%02x%c", byt, delim))
	}
	return buf.Bytes()
}
