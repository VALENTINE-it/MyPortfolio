package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net/http"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type result struct {
	duration time.Duration
	status   int
	err      error
}

func main() {
	url := flag.String("url", "http://localhost:8080/api/projects", "Target URL to test")
	concurrency := flag.Int("c", 100, "Number of concurrent virtual users")
	durationSec := flag.Int("d", 10, "Duration of the test in seconds")
	flag.Parse()

	fmt.Printf("Starting load test against %s\nConcurrency: %d workers | Duration: %ds\n", *url, *concurrency, *durationSec)

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 1000,
			IdleConnTimeout:     30 * time.Second,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
	}

	stop := make(chan struct{})
	var wg sync.WaitGroup
	var totalReqs int64
	var totalErrors int64
	resultsChan := make(chan result, 50000)

	// Collect results
	var allDurations []float64
	statusCodes := make(map[int]int)
	collectorDone := make(chan struct{})

	go func() {
		for res := range resultsChan {
			if res.err != nil || res.status >= 400 {
				atomic.AddInt64(&totalErrors, 1)
			}
			if res.status != 0 {
				statusCodes[res.status]++
			}
			if res.duration > 0 {
				allDurations = append(allDurations, float64(res.duration.Microseconds())/1000.0)
			}
		}
		close(collectorDone)
	}()

	start := time.Now()

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					reqStart := time.Now()
					req, err := http.NewRequest(http.MethodGet, *url, nil)
					if err != nil {
						resultsChan <- result{err: err}
						continue
					}
					req.Header.Set("User-Agent", "Go-LoadTest-Agent")

					resp, err := client.Do(req)
					reqDuration := time.Since(reqStart)

					if err != nil {
						resultsChan <- result{duration: reqDuration, err: err}
						continue
					}

					_, _ = io.Copy(io.Discard, resp.Body)
					_ = resp.Body.Close()

					atomic.AddInt64(&totalReqs, 1)
					resultsChan <- result{
						duration: reqDuration,
						status:   resp.StatusCode,
					}
				}
			}
		}()
	}

	time.Sleep(time.Duration(*durationSec) * time.Second)
	close(stop)
	wg.Wait()
	close(resultsChan)
	<-collectorDone

	totalElapsed := time.Since(start).Seconds()
	rps := float64(totalReqs) / totalElapsed

	sort.Float64s(allDurations)
	p50, p95, p99 := 0.0, 0.0, 0.0
	n := len(allDurations)
	if n > 0 {
		p50 = allDurations[int(float64(n)*0.50)]
		p95 = allDurations[int(float64(n)*0.95)]
		p99 = allDurations[int(float64(n)*0.99)]
	}

	fmt.Println("\n================ LOAD TEST RESULTS ================")
	fmt.Printf("Total Requests:       %d\n", totalReqs)
	fmt.Printf("Total Errors:         %d (%.2f%%)\n", totalErrors, float64(totalErrors)/float64(totalReqs+totalErrors)*100)
	fmt.Printf("Duration:             %.2fs\n", totalElapsed)
	fmt.Printf("Requests / Sec:       %.2f req/s\n", rps)
	fmt.Printf("Latency P50:          %.2f ms\n", p50)
	fmt.Printf("Latency P95:          %.2f ms\n", p95)
	fmt.Printf("Latency P99:          %.2f ms\n", p99)
	fmt.Println("Status Code Breakdown:")
	for code, count := range statusCodes {
		fmt.Printf("  HTTP %d: %d\n", code, count)
	}
	fmt.Println("====================================================")
}
