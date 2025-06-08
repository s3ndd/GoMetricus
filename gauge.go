package gometricus

import (
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
)

// statsdGauge implements the Gauge interface for statsd metrics.
type statsdGauge struct {
	name   string
	client *statsd.Client
	tags   []string
}

// Gauge creates a new Gauge instance with the given name.
func (m *statsdMetrics) Gauge(name string) Gauge {
	if name == "" {
		// Return a no-op gauge or handle error based on requirements
		panic(fmt.Errorf("gauge name cannot be empty"))
	}
	return &statsdGauge{
		name:   name,
		client: m.client,
		tags:   make([]string, 0),
	}
}

// Update sets the gauge to the specified value.
func (g *statsdGauge) Update(val int64) {
	if err := g.client.Gauge(g.name, float64(val), g.tags, DefaultRate); err != nil {
		// Log error if needed, depending on application requirements
	}
}

// WithTags returns a new Gauge with the provided tags appended.
func (g *statsdGauge) WithTags(tags []string) Gauge {
	newTags := make([]string, 0, len(g.tags)+len(tags))
	newTags = append(newTags, g.tags...)
	newTags = append(newTags, tags...)
	return &statsdGauge{
		name:   g.name,
		client: g.client,
		tags:   newTags,
	}
}

// WithTag returns a new Gauge with a single tag appended.
func (g *statsdGauge) WithTag(name, value string) Gauge {
	tag := fmt.Sprintf("%s:%s", name, value)
	return g.WithTags([]string{tag})
}

// Tags returns the current tags associated with the gauge.
func (g *statsdGauge) Tags() []string {
	return g.tags
}
