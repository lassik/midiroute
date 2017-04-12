package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

func stringIter(ss []string, i int) func() (string, bool) {
	return func() (string, bool) {
		if i < len(ss) {
			i++
			return ss[i-1], false
		} else {
			return "", true
		}
	}
}

var intoFormat = msgToMsg

func setup() error {
	if isTerminalFd(os.Stdout.Fd()) {
		intoFormat = msgToDump
	}
	readArg := stringIter(os.Args, 1)
	for arg, done := readArg(); !done; arg, done = readArg() {
		switch arg {
		default:
			return errors.New(fmt.Sprintf("cannot figure out what you mean by %q", arg))
		}
	}
	return nil
}

var msgBuf bytes.Buffer

func flushMsgBuf() {
	msg := msgBuf.Bytes()
	msgBuf.Reset()
	// TODO: Some formats may need to batch together multiple messages.
	// Then we'll have to use channels instead of simple function calls.
	msg = intoFormat(msg)
	os.Stdout.Write(msg)
}

func receiveMidiByte(c byte) {
	if isStatusByte(c) {
		flushMsgBuf()
	}
	msgBuf.WriteByte(c)
}
