package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	proxy_checker "github.com/trananhtung/proxy-checker"
)

func main() {
	// URL of the file to download
	url := "https://raw.githubusercontent.com/officialputuid/KangProxy/KangProxy/socks4/socks4.txt"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to send the HTTP GET request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download the file: %s\n", response.Status)
		return
	}

	// convert the response body to type io.Reader to string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read the response body: %v\n", err)
		return
	}

	// convert the body to type string
	sb := string(body)

	// split the body into lines
	proxies := strings.Split(sb, "\n")

	// check the availability of each proxy use goroutine
	var wg sync.WaitGroup
	// set max goroutine in waitgroup

	for _, proxy := range proxies {
		wg.Add(1)
		go func(proxy string) {
			defer wg.Done()
			httpProxy := "socks4://" + proxy
			if proxy_checker.ProxyTest(httpProxy, proxy_checker.AMAZON_CHECK_IP_URL, 5) {
				fmt.Printf("Proxy %s is available\n", httpProxy)
			}
		}(proxy)
	}
	wg.Wait()
}
