package main

import (
	"context"
	"fmt"
	"os"
)

func help() {
	fmt.Println("netspeed v0.1\n\nUsage:\n-d | --download Calculate average download speed\n-u | --upload Calculate average upload speed\n-b | --both Calculate both download and upload speed")
	return
}

func main() {
	ctx := context.Background()
	args := os.Args

	if len(args) == 1 {
		help()
		return
	}

	switch args[1] {
	case "-d", "--download":
		out, err := dlspeed(ctx)

		if err != nil {
			fmt.Println("Something went wrong.\nError log: " + err.Error())
			return
		}

		fmt.Println("\nAverage download speed: " + out)
		break

	case "-u", "--upload":
		out, err := ulspeed(ctx)

		if err != nil {
			fmt.Println("Something went wrong.\nError log: " + err.Error())
			return
		}

		fmt.Println("\nAverage upload speed: " + out)
		break

	case "-b", "--both":
		dl, err := dlspeed(ctx)

		if err != nil {
			fmt.Println("Something went wrong.\nError log: " + err.Error())
			return
		}

		// keep an empty newline for the second message
		fmt.Println()

		ul, err := ulspeed(ctx)

		if err != nil {
			fmt.Println("Something went wrong.\nError log: " + err.Error())
			return
		}

		// another new line to print the first result after testing log
		fmt.Println()

		fmt.Println("Average download speed: " + dl)
		fmt.Println("Average upload speed: " + ul)
		break

	default:
		help()
		break
	}
}
