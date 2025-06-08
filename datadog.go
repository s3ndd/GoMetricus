package gometricus

import (
	"fmt"
	"strings"

	"github.com/DataDog/datadog-go/statsd"
)

// statsdMetrics implements the MetricsInterface for Datadog statsd.
type statsdMetrics struct {
	client *statsd.Client
}

// SetupDatadog initializes a statsd-based MetricsInterface and sets it as the default metrics provider.
func SetupDatadog(appPrefix, addr string, tags []string) error {
	if appPrefix == "" {
		return fmt.Errorf("appPrefix cannot be empty")
	}
	if addr == "" {
		return fmt.Errorf("addr cannot be empty")
	}

	m, err := NewStatsdMetrics(appPrefix, addr, tags)
	if err != nil {
		return fmt.Errorf("failed to create statsd metrics: %w", err)
	}

	SetMetrics(m)
	return nil
}

// NewStatsdMetrics creates a new MetricsInterface instance for Datadog statsd.
func NewStatsdMetrics(appPrefix, addr string, tags []string) (MetricsInterface, error) {
	// Ensure appPrefix ends with a dot
	if !strings.HasSuffix(appPrefix, ".") {
		appPrefix += "."
	}

	client, err := statsd.New(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize statsd client: %w", err)
	}

	// Clean and normalize tags
	cleanedTags := make([]string, 0, len(tags))
	for _, tag := range tags {
		cleanedTag := strings.ToLower(strings.TrimSpace(tag))
		if cleanedTag != "" {
			cleanedTags = append(cleanedTags, cleanedTag)
		}
	}

	client.Namespace = appPrefix
	client.Tags = cleanedTags

	return &statsdMetrics{
		client: client,
	}, nil
}
