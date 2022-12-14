package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// private function, used to request to the server
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

// function to test the download speed
func dlspeed(ctx context.Context) (string, error) {
	exit := make(chan bool)
	url, err := getClosetServer(ctx)

	url = strings.ReplaceAll(url, "upload.php", "random1000x1000.jpg")

	if err != nil {
		return "", err
	}

	fmt.Print("Testing: ")

	// execute progressbar asynchronously
	go progressBar(exit)

	speed := 0.0
	startTime := time.Now()

	// maximum 6 rounds, increase according your needs
	for i := 0; i < 6; i++ {
		dl(ctx, url)

		endTime := time.Now()
		reqMB := 1000 * 1000 * 2 / 1000 / 1000

		speed = float64(reqMB) * 8 * float64(6) / endTime.Sub(startTime).Seconds()

	}
	v := fmt.Sprintf("%5.2fmbps", speed/8.0)
	exit <- true
	return v, nil
}
