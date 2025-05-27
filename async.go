package golibsasync

func Async(fn func()) {
	if fn == nil {
		return
	}

	// Create a new goroutine to execute the function asynchronously
	go func() {
		fn()
	}()
}
