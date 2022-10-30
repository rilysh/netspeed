package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Private function, used to request to the server
func dl(ctx context.Context, url string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		fmt.Println("An error occurred when trying to create a new http context.\nFull log: " + err.Error())
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("An error occurred when trying to send GET request to the server.\nFull log: " + err.Error())
		return
	}

	defer resp.Body.Close()

	_, err = io.Copy(io.Discard, resp.Body)

	if err != nil {
		fmt.Println("Failed to copy the buffer.\nFull log: " + err.Error())
		return
	}
}

// Function to test the download speed
func dlspeed(ctx context.Context) (string, error) {
	exit := make(chan bool)
	url, err := getClosetServer(ctx)

	url = strings.ReplaceAll(url, "upload.php", "random1000x1000.jpg")

	if err != nil {
		return "", err
	}
	// As it's a progress we'll execute it as gorotine
	fmt.Print("Testing: ")
	go progressBar(exit)

	average := 0.0
	// _, err := serversCount(ctx)

	if err != nil {
		// Should be handled by serversCount function
		return "", err
	}
	latency, err := checkLatency(ctx)

	if err != nil {
		return "", err
	}
	// First 10 servers
	for i := 0; i < 10; i++ {
		startTime := time.Now()
		dl(ctx, url)
		endTime := time.Now()
		waited := endTime.Sub(startTime.Add(latency)).Seconds()
		if waited < 0 {
			waited = endTime.Sub(startTime).Seconds()
		}
		speed := 1.125 * 8 * 2 / waited

		average += speed
	}
	v := fmt.Sprintf("%5.2fmbps", (average/10.0)/8)
	exit <- true
	return v, nil
}
