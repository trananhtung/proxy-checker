package proxy_checker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxyTest_Success(t *testing.T) {

	mockProxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Successful connection
	}))
	defer mockProxyServer.Close()

	proxy := mockProxyServer.URL
	urlTarget := AMAZON_CHECK_IP_URL

	result := ProxyTest(proxy, urlTarget, 5)

	if !result {
		t.Errorf("ProxyTest failed. Expected true, got false")
	}
}

func TestProxyTest_Fail(t *testing.T) {

	mockProxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer mockProxyServer.Close()

	proxy := mockProxyServer.URL
	urlTarget := AMAZON_CHECK_IP_URL

	result := ProxyTest(proxy, urlTarget, 5)

	if result {
		t.Errorf("ProxyTest failed. Expected false, got true")
	}
}

func TestProxyTest_InvalidProxy(t *testing.T) {

	proxy := "invalid_proxy"
	urlTarget := AMAZON_CHECK_IP_URL

	result := ProxyTest(proxy, urlTarget, 5)

	if result {
		t.Errorf("ProxyTest failed. Expected false, got true")
	}
}

func TestProxyTest_Timeout(t *testing.T) {

	mockProxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Successful connection
	}))
	defer mockProxyServer.Close()

	proxy := mockProxyServer.URL
	urlTarget := AMAZON_CHECK_IP_URL

	result := ProxyTest(proxy, urlTarget, 0)

	if result {
		t.Errorf("ProxyTest failed. Expected false, got true")
	}
}
