package gometricus

import (
	"log"
	"os"
	"sync/atomic"
)

var shared atomic.Value

func Metrics() MetricsInterface {
	if shared.Load() == nil {
		metricsPrefix := os.Getenv("METRICS_PREFIX")
		if metricsPrefix == "" {
			metricsPrefix = os.Getenv("SOURCE_PROGRAM")
			if metricsPrefix == "" {
				metricsPrefix = "unknown"
			}
		}

		statsdHost := os.Getenv("STATSD_HOST")
		if statsdHost == "" {
			statsdHost = "localhost:8125"
		}

		err := EnableDatadog(metricsPrefix, statsdHost, []string{})
		if err != nil {
			log.Fatalf("Could not initialize Datadog: %v", err)
		}
	}
	return shared.Load().(MetricsInterface)
}

func EnableDatadog(appPrefix, addr string, tags []string) error {
	return SetupDatadog(appPrefix, addr, tags)
}

func SetMetrics(metrics MetricsInterface) {
	shared.Store(metrics)
}
