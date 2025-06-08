# GoMetricus

Metrics for Go — lightweight, concurrent-safe, and with built-in support for exporting Counters, Gauges, and Timers to [DataDog](https://docs.datadoghq.com/developers/dogstatsd/) via DogStatsD.

![Go](https://img.shields.io/badge/Go-1.20+-blue)
[![MIT License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

---

## ✨ Introduction

When your system is running in production, it's critical to understand **how** it's running:

- Is it handling more traffic than usual?
- Are some parts of your application slowing down?
- Where is your system spending time?

**GoMetricus** is a small metrics library that helps you answer these questions by recording and exporting structured metrics.

---

## 🚀 Simple Example

Suppose you want to track how long it takes to insert customers into your database and how often it fails:

```go
var insertTimer = gometricus.Metrics().Timer("customer.insert")
var insertFailureCounter = gometricus.Metrics().Counter("customer.insert_failure")

func InsertCustomer(customer Customer) error {
    timer := insertTimer.Start()
    defer timer.Stop()

    err := db.Save(customer)
    if err != nil {
        insertFailureCounter.Inc(1)
        return err
    }
    return nil
}
```

## ⚙️ Initializing Metrics
Metrics are safe for concurrent use. Initialize them early, preferably during setup:
```go
type APIClient struct {
    requestTimer gometricus.Timer
}

func NewAPIClient() *APIClient {
    return &APIClient{
        requestTimer: gometricus.Metrics().Timer("api.request"),
    }
}
```

## 🏷 Using Tags
You can tag metrics with structured context (e.g., status_code:200, region:us-east-1) to filter and group them in dashboards:
```go
var responseCounter gometricus.Counter

func handleResponse(r *http.Response) {
    responseCounter.WithTag("status_code", fmt.Sprint(r.StatusCode)).Inc(1)

    responseCounter.WithTags([]string{
        "env:prod",
        "region:us-east-1",
    }).Inc(1)
}
```
> Note: Avoid using high-cardinality tags like user IDs. Keep tag values bounded (<100 distinct values ideally).

## 📊 Metric Types
### Counter
```go
var reqCounter gometricus.Counter

func Init() {
    reqCounter = gometricus.Metrics().Counter("requests.total")
}

func HandleRequest() {
    reqCounter.Inc(1)
}

```

### Gauge
```go
var queueGauge gometricus.Gauge

func StartQueueTracking() {
    queueGauge = gometricus.Metrics().Gauge("queue.size")

    go func() {
        ticker := time.NewTicker(time.Second)
        for range ticker.C {
            queueGauge.Update(int64(getQueueLength()))
        }
    }()
}

```

### Timer
```go
var dbTimer gometricus.Timer

func Init() {
    dbTimer = gometricus.Metrics().Timer("db.query")
}

func Query() {
    timer := dbTimer.Start()
    defer timer.Stop()

    db.ExecuteQuery()
}
```
Metrics emitted include:
- db.query.count
- db.query.avg
- db.query.max
- db.query.median
- db.query.95percentile

## 🌐 HTTP Timing Middleware
Wrap your HTTP handlers with built-in timers:
```go
http.Handle("/foo", gometricus.NewTimedHandler("http.foo", http.HandlerFunc(myHandler), nil))

http.HandleFunc("/bar", gometricus.NewTimedHandlerFunc("http.bar", func(w http.ResponseWriter, r *http.Request) {
    // handle request
}))

```

## 🐶 DataDog Integration
To enable DataDog output, set these environment variables:
```go
METRICS_PREFIX=your_service_name
STATSD_HOST=localhost:8125
```
Or enable it directly in code:
```go
err := gometricus.EnableDatadog(
    "your_service_name",
    "localhost:8125",   
    []string{"env:dev"}, 
)
if err != nil {
    log.Fatalf("Failed to enable metrics: %v", err)
}
```

## 📦 Installation
```shell
go get github.com/s3ndd/gometricus

```

