package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

const serverURL = "https://www.speedtest.net/api/js/servers?engine=js&limit=10"

type JSONResponse []struct {
	URL      string `json:"url"`
	Distance int    `json:"distance"`
}

// Get the first server
func getServer(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL, nil)

	if err != nil {
		fmt.Println("An error occurred when trying to create a new http context.\nFull log: " + err.Error())
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("An error occurred when trying to send GET request to the server.\nFull log: " + err.Error())
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Failed to read response body.\nFull log: " + err.Error())
		return "", err
	}

	var jsonResp JSONResponse
	json.Unmarshal([]byte(body), &jsonResp)

	if len(jsonResp) != 0 {
		return jsonResp[0].URL, nil
	} else {
		return "No servers found", nil
	}
}

// Get the closet server according to the distance
// Note: Here distance doesn't much matters.
// It's recommended to use the first server that the API returns
// as that's most likely fastet server than other ones
func getClosetServer(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL, nil)

	if err != nil {
		fmt.Println("An error occurred when trying to create a new http context.\nFull log: " + err.Error())
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("An error occurred when trying to send GET request to the server.\nFull log: " + err.Error())
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Failed to read response body.\nFull log: " + err.Error())
		return "", err
	}

	var jsonResp JSONResponse
	json.Unmarshal([]byte(body), &jsonResp)

	var distances []int
	for i := 0; i < len(jsonResp); i++ {
		distances = append(distances, int(jsonResp[i].Distance))
	}

	sort.Ints(distances)
	for i := 0; i < len(jsonResp); i++ {
		if jsonResp[i].Distance == distances[0] {
			return jsonResp[i].URL, nil
		}
	}
	return "", nil
}
