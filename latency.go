package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Check client to server latency
func checkLatency(ctx context.Context) (time.Duration, error) {
	url, err := getServer(ctx)
	if err != nil {
		// Should show the error caused by server implementation
		return -1, err
	}

	url = strings.ReplaceAll(url, "upload.php", "latency.txt")

	lowVal := time.Second * 10

	for i := 0; i < 3; i++ {
		startTime := time.Now()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

		if err != nil {
			fmt.Println("An error occurred when trying to create a new http context.\nFull log: " + err.Error())
			return -1, err
		}

		client := http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("An error occurred when trying to send GET request to the server.\nFull log: " + err.Error())
			return -1, err
		}

		endTime := time.Now()

		if endTime.Sub(startTime) < lowVal {
			lowVal = endTime.Sub(startTime)
		}

		resp.Body.Close()
	}

	latency := time.Duration(int64(lowVal.Nanoseconds() / 2))

	return latency, nil
}
