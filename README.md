# Proxy Checker
The proxy_checker package provides a library for checking the availability of proxies (HTTP, HTTPS, SOCKS4, SOCKS5).

## Installation
To use this library, you need to have Go installed and set up your Go workspace.

Run the following command to install the proxy_checker package:

```bash
go get github.com/trananhtung/proxy_checker
```

## Usage
To check the availability of a proxy, use the ProxyTest function. It performs a test request to the specified target URL using the provided proxy.

```go
import (
	"fmt"
	"github.com/trananhtung/proxy_checker"
)

func main() {
	proxy := "http://1.1.1.1:8080"
	targetURL := "http://example.com"
	timeout := uint(5) // Timeout in seconds

	result := proxy_checker.ProxyTest(proxy, targetURL, timeout)

	if result {
		fmt.Println("Proxy is available")
	} else {
		fmt.Println("Proxy is not available")
	}
}
```

## Function Documentation
### func ProxyTest

```go
func ProxyTest(proxy, urlTarget string, timeout uint) bool
```

ProxyTest checks the availability of a proxy by performing a test request.

- It sends an HTTP GET request to the specified target URL using the provided proxy.
- The function supports both HTTP and HTTPS proxies .
- It uses the specified timeout duration for the request.

The function returns true if the proxy is available and the test request succeeds. It returns false if the proxy is unavailable or the test request fails.

### Parameters:

- proxy: The proxy URL in the format http://<ip>:<port>, https://<ip>:<port>,  socks4://<ip>:<port>, socks5://<ip>:<port>.
- urlTarget: The target URL to send the test request to.
- timeout: The timeout duration in seconds for the test request.
