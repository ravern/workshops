package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var (
	rateLimit   int64 = 3
	currentRate int64 = 0
)

type video struct {
	title string
	size  int
}

func (v video) download() {
	atomic.AddInt64(&currentRate, 1)
	fmt.Printf("Downloading %s...\n", v.title)
	if currentRate > rateLimit {
		fmt.Println("Download rate limit exceeded! Please try again later.")
		os.Exit(1)
	}
	time.Sleep(time.Duration(v.size) * time.Millisecond)
	fmt.Printf("Finished downloading %s!\n", v.title)
	atomic.AddInt64(&currentRate, -1)
}

func downloadVideo(downloadQueue <-chan video) {
	for v := range downloadQueue {
		v.download()
	}
}

func main() {
	videos := []video{
		{"Go Concurrency Tutorial", 500},
		{"How to Build Microservices", 300},
		{"Understanding Goroutines", 700},
		{"Introduction to Go", 200},
		{"Go Best Practices", 400},
	}

	downloadQueue := make(chan video, len(videos))
	for i := 0; i < int(rateLimit); i++ {
		go downloadVideo(downloadQueue)
	}
	for _, v := range videos {
		downloadQueue <- v
	}

	time.Sleep(5 * time.Second)
}
