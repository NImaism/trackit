package network

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

func InitHTTPClient(timeoutSec int, verifySSL, disableRedirect bool) *http.Client {
	timeout := time.Duration(timeoutSec) * time.Second
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifySSL},
	}

	httpClient := http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	if disableRedirect {
		httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return &httpClient
}

func CheckValidURL(input string) bool {
	parsedURL, err := url.ParseRequestURI(input)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}
