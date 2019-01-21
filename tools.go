package utilz

// Days calculates total days given initial number and frequency
func Days(init, freq int) int {
	ret := 0
	articles := float64(init)
	for articles >= 1 {
		ret += int(articles)                // newly added articles which are actual
		articles = articles/float64(freq) + // newly added articles in the round, including actual ones and potential ones
			articles - float64(int(articles)) // potential articles in the last round
	}
	return ret
}
