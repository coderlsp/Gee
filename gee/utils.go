package gee

// Program will panic if guard is false.
func assert(guard bool, text string) {
	if !guard {
		panic(text)
	}
}
