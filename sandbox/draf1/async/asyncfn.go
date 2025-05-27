package async

import "github.com/fengdotdev/golibs-async/sandbox/draf1/eventloop"

var (
	el = eventloop.New()
)

func Async(fn func()) {
	if fn == nil {
		return
	}

	// Register the function to be executed in the event loop
	el.Async(fn)

}
