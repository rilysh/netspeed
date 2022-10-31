package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// private function to create upload request to the server
func ul(ctx context.Context, purl string) {
	val := url.Values{}

	val.Add("content", strings.Repeat("0123456789", 2800*100-51))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, purl, strings.NewReader(val.Encode()))

	if err != nil {
		fmt.Println("An error occurred when trying to create a new http context.\nFull log: " + err.Error())
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

// function to test the upload speed
func ulspeed(ctx context.Context) (string, error) {
	exit := make(chan bool)
	url, err := getClosetServer(ctx)

	if err != nil {
		return "", err
	}

	fmt.Print("Testing: ")
	// execute progressbar asynchronously
	go progressBar(exit)

	speed := 0.0
	startTime := time.Now()

	// maximum 6 rounds, increasde according your needs
	for i := 0; i < 6; i++ {
		ul(ctx, url)

		endTime := time.Now()
		reqMB := 3

		speed = float64(reqMB) * 8 * float64(6) / endTime.Sub(startTime).Seconds()

	}
	v := fmt.Sprintf("%5.2fmbps", speed/8.0)
	exit <- true
	return v, nil
}
