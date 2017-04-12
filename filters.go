package main

type filterFunc func(mIn chan []byte, mOut chan []byte)

func identityFilter(args []string) filterFunc {
	return func(mIn chan []byte, mOut chan []byte) {
		for m := range mIn {
			mOut <- m
		}
	}
}

func transposeFilter(args []string) filterFunc {
	const semitones int = 1
	return func(mIn chan []byte, mOut chan []byte) {
		for m := range mIn {
			if isMsg(m, NoteOn) || isMsg(m, NoteOff) {
				m[1] = clampToByte(int(m[1]) + semitones)
			}
			mOut <- m
		}
	}
}
