package async

import (
	"time"

	"github.com/fengdotdev/golibs-async/sandbox/draf1/eventloop"
	"github.com/fengdotdev/golibs-async/sandbox/draf1/future"
)

var (
	el = eventloop.New()
)

func Async(fn interface{}) {
	time.Sleep(100 * time.Millisecond) // Simulate some delay

	if f, ok := fn.(func()); ok {
		// If fn is a function, call it directly
		f()
		return
	}
	if f, ok := fn.(func() any); ok {
		// If fn is a function that returns a value, call it and ignore the return value
		f()
		return
	}

}

func Await[T any](fn func() T) *future.Future[T] {
	if fn == nil {
		panic("fn cannot be nil")
	}

	// Create a new future
	f := future.NewFuture[T]()

	// Use the event loop to run the function asynchronously
	el.Await(func() {
		t := fn()

		f.Resolve(t) // Resolve the future with the result of the function
	})

	return f
}
