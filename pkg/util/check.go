package util

// Check checks an error variable. If it is NOT nil, panic it.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
