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

	url := "https://kdkekd.requestcatcher.com/" // Replace with your API endpoint

	proxy := proxyList[rand.Intn(len(proxyList))]

	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")

	req.SetRequestURI(url)

	err := fasthttp.Do(req, nil)

	if err != nil {
		color.Println("[<fg=blue>"+host+"</>] - [<fg=red>"+url+"</>] - [<fg=red>Invalid Proxy</>] ::", proxy)

	} else {

		color.Println("[<fg=blue>"+host+"</>] - [<fg=green>"+url+"</>] - [<fg=green>200</>] ::", proxy)

	}

}
func Laddkod() {

	url := "https://kdkekd.requestcatcher.com/" // Replace with your API endpoint

	proxy := proxyList[rand.Intn(len(proxyList))]

	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	req.SetRequestURI(url)

	err := fasthttp.Do(req, nil)

	if err != nil {
		color.Println("[<fg=blue>"+host+"</>] - [<fg=red>"+url+"</>] - [<fg=red>Invalid Proxy</>] ::", proxy)

	} else {

		color.Println("[<fg=blue>"+host+"</>] - [<fg=green>"+url+"</>] - [<fg=green>200</>] ::", proxy)

	}

}
func ProdMobil2() {

	url := "https://kdkekd.requestcatcher.com/" // Replace with your API endpoint

	proxy := proxyList[rand.Intn(len(proxyList))]

	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod("GET")
	req.SetRequestURI(url)

	err := fasthttp.Do(req, nil)

	if err != nil {
		color.Println("[<fg=blue>"+host+"</>] - [<fg=red>"+url+"</>] - [<fg=red>Invalid Proxy</>] ::", proxy)

	} else {

		color.Println("[<fg=blue>"+host+"</>] - [<fg=green>"+url+"</>] - [<fg=green>200</>] ::", proxy)

	}

}
