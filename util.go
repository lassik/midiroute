package main

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
