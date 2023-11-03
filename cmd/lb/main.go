package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/LukaMacharashvili/Load-Balancer/internal/utils/slices"
	"github.com/LukaMacharashvili/Load-Balancer/internal/utils/urls"
)

func healthCheck(targetUrlString string) bool {
	targetUrlString = targetUrlString + "/health"

	resp, err := http.Get(targetUrlString)

	if err != nil || resp.StatusCode != 200 {
		return false
	}

	return true
}

func healthCheckScheduler(interval time.Duration, fullTargetUrls *[]string, targetUrls *[]string) {
	go func() {
		for {
			for _, targetUrl := range *fullTargetUrls {
				fmt.Printf("Health check to %s\n", targetUrl)
				if !healthCheck(targetUrl) {
					slices.RemoveElementString(targetUrls, targetUrl)
				} else {
					slices.AddTargetUrlToEnv(targetUrls, targetUrl)
				}
			}
			time.Sleep(interval)
		}
	}()
}

var (
	mu    sync.Mutex
	count int
)

func main() {
	mu.Lock()
	defer mu.Unlock()

	targetUrls := urls.GetTargetUrlsFromEnv()
	fullTargetUrls := targetUrls

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request to %s\n", targetUrls[count])
		proxy := urls.GetProxyFromUrlString(targetUrls[count])
		count = (count + 1) % len(targetUrls)
		proxy.ServeHTTP(w, r)
	})

	interval := 10 * time.Second
	healthCheckScheduler(interval, &fullTargetUrls, &targetUrls)

	http.ListenAndServe(":3002", nil)
}
