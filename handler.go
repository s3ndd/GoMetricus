package gometricus

import (
	"fmt"
	"net/http"
)

// TimedHandler wraps an http.Handler to track request timing metrics.
type TimedHandler struct {
	Handler http.Handler
	Timer   Timer
}

// NewTimedHandler creates a new TimedHandler that wraps the provided handler with a timer metric.
func NewTimedHandler(name string, handler http.Handler, tags []string) *TimedHandler {
	if name == "" {
		// Return a no-op handler or handle error based on requirements
		panic(fmt.Errorf("handler name cannot be empty"))
	}

	timer := Metrics().Timer(name).WithTags(tags)
	return &TimedHandler{
		Handler: handler,
		Timer:   timer,
	}
}

// NewTimedHandlerFunc creates a new http.HandlerFunc that wraps the provided handler function with a timer metric.
func NewTimedHandlerFunc(name string, handler http.HandlerFunc, tags []string) http.HandlerFunc {
	if name == "" {
		// Return a no-op handler or handle error based on requirements
		panic(fmt.Errorf("handler name cannot be empty"))
	}

	timer := Metrics().Timer(name).WithTags(tags)
	return func(w http.ResponseWriter, r *http.Request) {
		t := timer.Start()
		defer t.Stop()
		handler(w, r)
	}
}

// ServeHTTP executes the underlying handler while tracking request timing.
func (h *TimedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := h.Timer.Start()
	defer t.Stop()
	h.Handler.ServeHTTP(w, r)
}
