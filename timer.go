package gometricus

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

// statsdTimer implements the Timer interface for statsd metrics.
type statsdTimer struct {
	name   string
	client *statsd.Client
	tags   []string
}

// startedStatsdTimer tracks the start time for a timer and its underlying timer.
type startedStatsdTimer struct {
	startedTime time.Time
	underlying  *statsdTimer
}

// Timer creates a new Timer instance with the given name.
func (m *statsdMetrics) Timer(name string) Timer {
	if name == "" {
		// Return a no-op timer or handle error based on requirements
		panic(fmt.Errorf("timer name cannot be empty"))
	}
	return &statsdTimer{
		name:   name,
		client: m.client,
		tags:   make([]string, 0),
	}
}

// Start begins a new timer and returns a StartedTimer to stop it.
func (t *statsdTimer) Start() StartedTimer {
	return &startedStatsdTimer{
		startedTime: time.Now(),
		underlying:  t,
	}
}

// Update records the specified duration for the timer.
func (t *statsdTimer) Update(d time.Duration) {
	if err := t.client.Timing(t.name, d, t.tags, DefaultRate); err != nil {
		// Log error if needed, depending on application requirements
	}
}

// UpdateSince records the duration since the specified time.
func (t *statsdTimer) UpdateSince(from time.Time) {
	t.Update(time.Since(from))
}

// WithTags returns a new Timer with the provided tags appended.
func (t *statsdTimer) WithTags(tags []string) Timer {
	newTags := make([]string, 0, len(t.tags)+len(tags))
	newTags = append(newTags, t.tags...)
	newTags = append(newTags, tags...)
	return &statsdTimer{
		name:   t.name,
		client: t.client,
		tags:   newTags,
	}
}

// WithTag returns a new Timer with a single tag appended.
func (t *statsdTimer) WithTag(name, value string) Timer {
	tag := fmt.Sprintf("%s:%s", name, value)
	return t.WithTags([]string{tag})
}

// Tags returns the current tags associated with the timer.
func (t *statsdTimer) Tags() []string {
	return t.tags
}

// Stop records the duration since the timer was started.
func (st *startedStatsdTimer) Stop() {
	st.underlying.UpdateSince(st.startedTime)
}
