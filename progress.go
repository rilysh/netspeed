package main

import (
	"fmt"
	"time"
)

// Progress bar (credit: @showwin)
func progressBar(exit chan bool) {
	for {
		select {
		case <-exit:
			break
		default:
			fmt.Print("#")
			time.Sleep(time.Second) // 1 second each time
		}
	}
}
