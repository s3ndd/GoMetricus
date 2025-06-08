package gometricus

import (
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
)

// DefaultRate is the default sampling rate for metrics.
const DefaultRate = 1.0

// statsdCounter implements the Counter interface for statsd metrics.
type statsdCounter struct {
	name   string
	client *statsd.Client
	tags   []string
}

// Counter creates a new Counter instance with the given name.
func (m *statsdMetrics) Counter(name string) Counter {
	if name == "" {
		// Return a no-op counter or handle error based on requirements
		panic(fmt.Errorf("counter name cannot be empty"))
	}
	return &statsdCounter{
		name:   name,
		client: m.client,
		tags:   make([]string, 0),
	}
}

// Inc increments the counter by the specified value.
func (c *statsdCounter) Inc(val int64) {
	if err := c.client.Count(c.name, val, c.tags, DefaultRate); err != nil {
		// Log error if needed, depending on application requirements
	}
}

// Dec decrements the counter by the specified value.
func (c *statsdCounter) Dec(val int64) {
	if err := c.client.Count(c.name, -val, c.tags, DefaultRate); err != nil {
		// Log error if needed, depending on application requirements
	}
}

// WithTags returns a new Counter with the provided tags appended.
func (c *statsdCounter) WithTags(tags []string) Counter {
	newTags := make([]string, 0, len(c.tags)+len(tags))
	newTags = append(newTags, c.tags...)
	newTags = append(newTags, tags...)
	return &statsdCounter{
		name:   c.name,
		client: c.client,
		tags:   newTags,
	}
}

// WithTag returns a new Counter with a single tag appended.
func (c *statsdCounter) WithTag(name, value string) Counter {
	tag := fmt.Sprintf("%s:%s", name, value)
	return c.WithTags([]string{tag})
}

// Tags returns the current tags associated with the counter.
func (c *statsdCounter) Tags() []string {
	return c.tags
}
