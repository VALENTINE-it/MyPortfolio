package monitoring

import (
	"fmt"
	"net/http"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// Metrics holds concurrency-safe application runtime metrics
type Metrics struct {
	mu sync.RWMutex

	// Counters
	totalRequests  int64
	activeRequests int64
	totalErrors    int64
	rateLimitedReq int64

	cacheHits   int64
	cacheMisses int64

	dbQueriesTotal int64
	dbErrorsTotal  int64
	dbLatencySumMs int64

	startTime time.Time

	// Route & status counts
	statusCounts map[int]int64
	routeCounts  map[string]int64

	// Duration tracking (latency buckets for p50, p95, p99)
	latenciesMs []float64
}

var (
	globalMetrics     *Metrics
	globalMetricsOnce sync.Once
)

// GetMetrics returns the singleton metrics collector
func GetMetrics() *Metrics {
	globalMetricsOnce.Do(func() {
		globalMetrics = &Metrics{
			startTime:    time.Now(),
			statusCounts: make(map[int]int64),
			routeCounts:  make(map[string]int64),
			latenciesMs:  make([]float64, 0, 1000),
		}
	})
	return globalMetrics
}

// IncActiveRequests increments active requests
func (m *Metrics) IncActiveRequests() {
	atomic.AddInt64(&m.activeRequests, 1)
}

// DecActiveRequests decrements active requests
func (m *Metrics) DecActiveRequests() {
	atomic.AddInt64(&m.activeRequests, -1)
}

// RecordRequest records a completed HTTP request
func (m *Metrics) RecordRequest(method, path string, status int, duration time.Duration) {
	atomic.AddInt64(&m.totalRequests, 1)
	if status >= 400 {
		atomic.AddInt64(&m.totalErrors, 1)
	}

	durationMs := float64(duration.Microseconds()) / 1000.0

	m.mu.Lock()
	defer m.mu.Unlock()

	m.statusCounts[status]++
	m.routeCounts[method+" "+path]++

	// Keep rolling sample of recent 2000 latencies
	if len(m.latenciesMs) >= 2000 {
		m.latenciesMs = m.latenciesMs[500:] // prune oldest
	}
	m.latenciesMs = append(m.latenciesMs, durationMs)
}

// RecordRateLimitRejected records a 429 rejection
func (m *Metrics) RecordRateLimitRejected() {
	atomic.AddInt64(&m.rateLimitedReq, 1)
}

// RecordCacheHit records cache hit
func (m *Metrics) RecordCacheHit() {
	atomic.AddInt64(&m.cacheHits, 1)
}

// RecordCacheMiss records cache miss
func (m *Metrics) RecordCacheMiss() {
	atomic.AddInt64(&m.cacheMisses, 1)
}

// RecordDBQuery records database query duration and error status
func (m *Metrics) RecordDBQuery(duration time.Duration, err error) {
	atomic.AddInt64(&m.dbQueriesTotal, 1)
	atomic.AddInt64(&m.dbLatencySumMs, duration.Milliseconds())
	if err != nil {
		atomic.AddInt64(&m.dbErrorsTotal, 1)
	}
}

// PrometheusHandler exposes metrics in standard Prometheus exposition format
func (m *Metrics) PrometheusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.mu.RLock()
		statusCopy := make(map[int]int64, len(m.statusCounts))
		for k, v := range m.statusCounts {
			statusCopy[k] = v
		}

		latenciesCopy := make([]float64, len(m.latenciesMs))
		copy(latenciesCopy, m.latenciesMs)
		m.mu.RUnlock()

		// Calculate percentiles
		p50, p95, p99 := calculatePercentiles(latenciesCopy)

		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		uptimeSeconds := time.Since(m.startTime).Seconds()

		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")

		fmt.Fprintf(w, "# HELP http_requests_total Total number of HTTP requests\n")
		fmt.Fprintf(w, "# TYPE http_requests_total counter\n")
		fmt.Fprintf(w, "http_requests_total %d\n\n", atomic.LoadInt64(&m.totalRequests))

		fmt.Fprintf(w, "# HELP http_active_requests Currently executing HTTP requests\n")
		fmt.Fprintf(w, "# TYPE http_active_requests gauge\n")
		fmt.Fprintf(w, "http_active_requests %d\n\n", atomic.LoadInt64(&m.activeRequests))

		fmt.Fprintf(w, "# HELP http_errors_total Total number of HTTP requests with status >= 400\n")
		fmt.Fprintf(w, "# TYPE http_errors_total counter\n")
		fmt.Fprintf(w, "http_errors_total %d\n\n", atomic.LoadInt64(&m.totalErrors))

		fmt.Fprintf(w, "# HELP rate_limit_rejections_total Total rate limit rejections\n")
		fmt.Fprintf(w, "# TYPE rate_limit_rejections_total counter\n")
		fmt.Fprintf(w, "rate_limit_rejections_total %d\n\n", atomic.LoadInt64(&m.rateLimitedReq))

		fmt.Fprintf(w, "# HELP cache_hits_total Total cache hits\n")
		fmt.Fprintf(w, "# TYPE cache_hits_total counter\n")
		fmt.Fprintf(w, "cache_hits_total %d\n\n", atomic.LoadInt64(&m.cacheHits))

		fmt.Fprintf(w, "# HELP cache_misses_total Total cache misses\n")
		fmt.Fprintf(w, "# TYPE cache_misses_total counter\n")
		fmt.Fprintf(w, "cache_misses_total %d\n\n", atomic.LoadInt64(&m.cacheMisses))

		fmt.Fprintf(w, "# HELP db_queries_total Total database queries\n")
		fmt.Fprintf(w, "# TYPE db_queries_total counter\n")
		fmt.Fprintf(w, "db_queries_total %d\n\n", atomic.LoadInt64(&m.dbQueriesTotal))

		fmt.Fprintf(w, "# HELP db_errors_total Total database errors\n")
		fmt.Fprintf(w, "# TYPE db_errors_total counter\n")
		fmt.Fprintf(w, "db_errors_total %d\n\n", atomic.LoadInt64(&m.dbErrorsTotal))

		fmt.Fprintf(w, "# HELP http_request_duration_ms HTTP request latency percentiles\n")
		fmt.Fprintf(w, "# TYPE http_request_duration_ms gauge\n")
		fmt.Fprintf(w, "http_request_duration_ms{quantile=\"0.50\"} %.2f\n", p50)
		fmt.Fprintf(w, "http_request_duration_ms{quantile=\"0.95\"} %.2f\n", p95)
		fmt.Fprintf(w, "http_request_duration_ms{quantile=\"0.99\"} %.2f\n\n", p99)

		fmt.Fprintf(w, "# HELP process_uptime_seconds Process uptime in seconds\n")
		fmt.Fprintf(w, "# TYPE process_uptime_seconds gauge\n")
		fmt.Fprintf(w, "process_uptime_seconds %.1f\n\n", uptimeSeconds)

		fmt.Fprintf(w, "# HELP go_goroutines Number of running goroutines\n")
		fmt.Fprintf(w, "# TYPE go_goroutines gauge\n")
		fmt.Fprintf(w, "go_goroutines %d\n\n", runtime.NumGoroutine())

		fmt.Fprintf(w, "# HELP go_memstats_alloc_bytes Number of allocated memory bytes\n")
		fmt.Fprintf(w, "# TYPE go_memstats_alloc_bytes gauge\n")
		fmt.Fprintf(w, "go_memstats_alloc_bytes %d\n\n", mem.Alloc)

		fmt.Fprintf(w, "# HELP go_memstats_sys_bytes Number of system memory bytes\n")
		fmt.Fprintf(w, "# TYPE go_memstats_sys_bytes gauge\n")
		fmt.Fprintf(w, "go_memstats_sys_bytes %d\n", mem.Sys)
	}
}

func calculatePercentiles(latencies []float64) (p50, p95, p99 float64) {
	n := len(latencies)
	if n == 0 {
		return 0, 0, 0
	}
	sort.Float64s(latencies)
	p50 = latencies[int(float64(n)*0.50)]
	p95 = latencies[int(float64(n)*0.95)]
	p99 = latencies[int(float64(n)*0.99)]
	return
}
