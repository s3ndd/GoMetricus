package gometricus

import "time"

type MetricsInterface interface {
	Timer(name string) Timer
	Gauge(name string) Gauge
	Counter(name string) Counter
	Meter(name string) Meter
}

type Timer interface {
	Start() StartedTimer
	Update(time.Duration)
	UpdateSince(time.Time)
	WithTags(tags []string) Timer
	WithTag(name, value string) Timer
	Tags() []string
}

type StartedTimer interface {
	Stop()
}

type Gauge interface {
	WithTags(tags []string) Gauge
	WithTag(name, value string) Gauge
	Update(int64)
	Tags() []string
}

type Meter interface {
	WithTags(tags []string) Meter
	WithTag(name, value string) Meter
	Mark(int64)
	Tags() []string
}

type Counter interface {
	WithTags(tags []string) Counter
	WithTag(name, value string) Counter
	Dec(int64)
	Inc(int64)
	Tags() []string
}
