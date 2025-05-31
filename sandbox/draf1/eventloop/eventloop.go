package eventloop

// avoid channels and goroutines target is single-threaded
// NO concurrency, no channels, no goroutines

type EventLoop struct {
	ratio             float64                // A ratio to control the processing of events
	priorityEvents    []func()               // A slice to hold event functions
	awaitEvents       []func()               // A slice to hold events that are waiting to be processed
	endProgram        bool                   // A flag to indicate if the event loop should stop
	mapresults        map[string]interface{} // A map to hold results of events
	evenLoopIsRunning bool                   // A flag to indicate if the event loop is currently running
}

func New() *EventLoop {

	el :=
		&EventLoop{
			priorityEvents: make([]func(), 0),
			awaitEvents:    make([]func(), 0),
			endProgram:     false,
			mapresults:     make(map[string]interface{}),
		}

	el.eventloop()
	return el
}

func (el *EventLoop) Run() {

}

func (el *EventLoop) eventloop() {
	el.evenLoopIsRunning = true
	// Process priority events first
	el.loopEventAwait()

	el.evenLoopIsRunning = false
}

func (el *EventLoop) loopEventPriority() {
	// Process priority events
	for _, event := range el.priorityEvents {
		event()
	}
	el.priorityEvents = el.priorityEvents[:0] // Clear the priority events after processing
}

func (el *EventLoop) loopEventAwait() {
	// Process await events
	for _, event := range el.awaitEvents {
		event()
	}
	el.awaitEvents = el.awaitEvents[:0] // Clear the await events after processing
}

func (el *EventLoop) runEventLoop() {
	if !el.evenLoopIsRunning {
		el.eventloop()
	}

}

func (el *EventLoop) Stop() {
	el.endProgram = true
}

func (el *EventLoop) Await(fn func()) {

	el.awaitEvents = append(el.awaitEvents, fn)

	el.runEventLoop()
}
