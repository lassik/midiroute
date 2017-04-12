package main

import (
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

func setup() error {
	readArg := stringIter(os.Args, 1)
	for arg, done := readArg(); !done; arg, done = readArg() {
		switch arg {
		default:
			return errors.New(fmt.Sprintf("cannot figure out what you mean by %q", arg))
		}
	}
	return nil
}

func receiveMidiByte(c byte) {
	fmt.Printf("CORE got midi byte 0x%02x\n", c)
}
