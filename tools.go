package tools

// Days calculates total days given initial number and frequency
func Days(init, freq int) int {
	ret := 0
	remain := init
	for remain != 0 {
		ret += remain
		remain /= freq
	}
	return ret
}
