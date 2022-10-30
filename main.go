package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	if len(os.Args) == 1 {
		fmt.Println("netspeed v0.1\n\nUsage: netspeed -d Calculate average download speed")
		return
	}

	if os.Args[1] == "-d" {
		out, err := dlspeed(ctx)

		if err != nil {
			panic(err) // This time we'd panic, as all other errors should be handled by given function
		}
		fmt.Println("\nAverage download speed: " + out)
	}
}
