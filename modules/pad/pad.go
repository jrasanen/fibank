package pad

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return out
}
func Right(str string, length int, pad string) string {
	return str + times(pad, length)
}
