package gometricus

import (
	"fmt"
)

// deprecatedMeter implements the Meter interface using a Counter, marked as deprecated.
type deprecatedMeter struct {
	name    string
	counter Counter
}

// Meter creates a new Meter instance with the given name, wrapping a Counter.
// Note: This implementation is deprecated; use Counter directly instead.
func (m *statsdMetrics) Meter(name string) Meter {
	if name == "" {
		// Return a no-op meter or handle error based on requirements
		panic(fmt.Errorf("meter name cannot be empty"))
	}
	return &deprecatedMeter{
		name:    name,
		counter: m.Counter(name),
	}
}

// Mark increments the meter by the specified value.
func (m *deprecatedMeter) Mark(val int64) {
	m.counter.Inc(val)
}

// WithTags returns a new Meter with the provided tags appended.
func (m *deprecatedMeter) WithTags(tags []string) Meter {
	return &deprecatedMeter{
		name:    m.name,
		counter: m.counter.WithTags(tags),
	}
}

// WithTag returns a new Meter with a single tag appended.
func (m *deprecatedMeter) WithTag(name, value string) Meter {
	tag := fmt.Sprintf("%s:%s", name, value)
	return m.WithTags([]string{tag})
}

// Tags returns the current tags associated with the meter.
func (m *deprecatedMeter) Tags() []string {
	return m.counter.Tags()
}
