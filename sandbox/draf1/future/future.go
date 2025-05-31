package future

import "fmt"

type Future[T any] struct {
	value     T
	completed bool
	then      func(T)
	onError   func(error)
	finally   func()
}

func NewFuture[T any]() *Future[T] {
	return &Future[T]{
		value:     *new(T), // Initialize with zero value of T
		completed: false,
		then:      func(T) {},
		onError:   func(error) {},
		finally:   func() {},
	}
}

func (f *Future[T]) IsCompleted() bool {
	return f.completed
}

func (f *Future[T]) String() string {
	if f.completed {
		return "Future completed with value: " + fmt.Sprintf("%v", f.value)
	}
	return "Future not completed"
}

func (f *Future[T]) Then(callback func(T)) *Future[T] {
	f.then = callback
	return f
}
func (f *Future[T]) Catch(callback func(error)) *Future[T] {
	f.onError = callback
	return f
}
func (f *Future[T]) Finally(callback func()) *Future[T] {
	f.finally = callback
	return f
}

func (f *Future[T]) Resolve(value T) {
	f.value = value
	f.completed = true
	if f.then != nil {
		f.then(value)
	}
	if f.finally != nil {
		f.finally()
	}
}
func (f *Future[T]) Reject(err error) {
	f.completed = true
	if f.onError != nil {
		f.onError(err)
	}
	if f.finally != nil {
		f.finally()
	}
}
