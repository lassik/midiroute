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

var fromChan = make(chan []byte)
var intoChan = make(chan []byte)
var intoFormat = msgToMsg

func setup() error {
	if isTerminalFd(os.Stdout.Fd()) {
		intoFormat = msgToDump
	}
	noArgsArray := [0]string{}
	noArgs := noArgsArray[:]
	var filter = identityFilter(noArgs)
	readArg := stringIter(os.Args, 1)
	for arg, done := readArg(); !done; arg, done = readArg() {
		switch arg {
		case "transpose":
			filter = transposeFilter(noArgs)
		default:
			return errors.New(fmt.Sprintf("cannot figure out what you mean by %q", arg))
		}
	}
	go filter(fromChan, intoChan)
	go output()
	return nil
}

var msgBuf bytes.Buffer

func flushMsgBuf() {
	msg := msgBuf.Bytes()
	msgBuf.Reset()
	fromChan <- msg
}

func output() {
	// TODO: Some formats may need to batch together multiple messages.
	// Then we'll have to use channels instead of simple function calls.
	for msg := range intoChan {
		os.Stdout.Write(intoFormat(msg))
	}
}

func receiveMidiByte(c byte) {
	if isStatusByte(c) {
		flushMsgBuf()
	}
	msgBuf.WriteByte(c)
}
