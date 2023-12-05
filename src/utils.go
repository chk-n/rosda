package internal

func prependBytes(s []byte, v byte) []byte {
	s = append(s, v)
	copy(s[1:], s)
	s[0] = v
	return s
}
